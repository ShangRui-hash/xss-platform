<template>
  <div>
    <el-button type="primary" size="small" @click="onClickCreateProjectBtn">创建项目</el-button>
    <el-table :data="projects" style="width: 100% border:0px">
      <el-table-column prop="id" label="#" width="50" type="index"></el-table-column>
      <el-table-column prop="name" label="项目名称" width="100">
        <template slot-scope="scope">
          <el-link type="warning" @click="onClickProjectName(scope.row)">{{scope.row.name}}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="desc" label="项目描述"></el-table-column>
      <el-table-column prop="loot_count" label="战利品数"></el-table-column>
      <el-table-column prop="created_at" label="创建时间"></el-table-column>
      <el-table-column prop="updated_at" label="更新时间"></el-table-column>
      <el-table-column fixed="right" label="操作" width="250">
        <template slot-scope="scope">
          <el-button @click="onClickUsageBtn(scope.row)" type="primary" size="small" plain>用 法</el-button>
          <el-button @click="onClickConfigBtn(scope.row)" type="warning" size="small" plain>配 置</el-button>
          <span style="margin-right:10px"></span>
          <el-popconfirm title="您确定要删除吗" @confirm="deleteProject(scope.row.id)">
            <el-button slot="reference" type="danger" size="small" plain>删除</el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- 创建项目的对话框 -->
    <el-dialog title="创建项目" :visible.sync="is_show_create_dialog">
      <ProjectForm :form="form"></ProjectForm>
      <div slot="footer" class="dialog-footer">
        <el-button @click="is_show_create_dialog=false">取消</el-button>
        <el-button type="primary" @click="createProject">确认</el-button>
      </div>
    </el-dialog>
    <!-- 配置项目的对话框 -->
    <el-dialog title="配置项目" :visible.sync="is_show_update_dialog">
      <ProjectForm :form="form"></ProjectForm>
      <div slot="footer" class="dialog-footer">
        <el-button @click="is_show_update_dialog=false">取消</el-button>
        <el-button type="primary" @click="updateProject">确认</el-button>
      </div>
    </el-dialog>
    <!-- 项目的使用说明对话框 -->
    <el-dialog title="使用说明" :visible.sync="is_show_usage_dialog">
      <usage :current_project="current_project"></usage>
    </el-dialog>
    <!-- 查看项目收到的战利品 -->
    <LootDrawer
      :project_name="current_project.name"
      :is_show_drawer="is_show_drawer"
      :loots="loots"
      @drawerClose="is_show_drawer = false"
    ></LootDrawer>
    <load-more-btn :loading="loading" :loadmore_btn_text="loadmore_btn_text" @loadmore="loadmore"></load-more-btn>
  </div>
</template>

<script>
import Usage from "@/components/Usage";
import ProjectForm from "@/components/ProjectForm";
import LootDrawer from "@/components/LootDrawer";
import LoadMoreBtn from "@/components/LoadMoreBtn";
export default {
  components: {
    Usage,
    ProjectForm,
    LootDrawer,
    LoadMoreBtn
  },
  data: () => {
    return {
      is_show_create_dialog: false,
      is_show_update_dialog: false,
      is_show_usage_dialog: false,
      is_show_drawer: false,
      form: {
        url_key: "",
        name: "",
        desc: "",
        module_list: []
      },
      projects: [],
      current_project: {},
      loots: [],
      offset: 0,
      count: 10,
      loadmore_btn_text: "",
      loading: false
    };
  },
  created() {
    this.loadmore_btn_text = this.config.loadmore_text;
    //请求所有项目
    this.getProjects();
  },
  methods: {
    //加载更多
    loadmore() {
      this.loading = true;
      this.getProjects();
    },
    //点击用法按钮
    onClickUsageBtn(row) {
      this.is_show_usage_dialog = true;
      this.current_project = row;
    },
    //当点击创建项目按钮时
    onClickCreateProjectBtn() {
      //1.显示模态框
      this.is_show_create_dialog = true;
      //2.获取表单格式
      this.getProjectForm();
    },
    //当点击项目名称时
    onClickProjectName(row) {
      this.is_show_drawer = true;
      this.current_project = row;
      this.$axios({
        method: "GET",
        url: `/api/v1/loots/${row.url_key}`
      })
        .then(resp => {
          this.loots = resp.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    //点击配置按钮
    onClickConfigBtn(row) {
      this.current_project = row;
      //1.显示模态框
      this.is_show_update_dialog = true;
      //2.查询项目详情
      this.getProjectDetail(row.id);
    },
    //获取项目详情
    getProjectDetail(id) {
      this.$axios({
        method: "GET",
        url: `/api/v1/projectform?id=${id}`
      })
        .then(resp => {
          this.form = resp.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    //获取项目列表
    getProjects() {
      this.$axios({
        method: "GET",
        url: `/api/v1/projects/${this.offset}/${this.count}`
      })
        .then(resp => {
          let projects = resp.data.data == null ? [] : resp.data.data;
          for (let i = 0; i < projects.length; i++) {
            let index = this.projects.findIndex(item => {
              return item.id == projects[i].id;
            });
            if (-1 == index) {
              this.projects.push(projects[i]);
            } else {
              this.projects[index] = projects[i];
            }
          }
          this.loading = false;
          this.offset = this.projects.length;
          if (projects.length < this.count) {
            this.loadmore_btn_text = this.config.nomore_text;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    //获取项目表单格式
    getProjectForm() {
      this.$axios({
        method: "GET",
        url: `/api/v1/projectform`
      })
        .then(resp => {
          this.form = resp.data.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    //更新项目
    updateProject() {
      this.$axios({
        method: "PUT",
        url: `/api/v1/project/${this.current_project.url_key}`,
        data: this.form
      })
        .then(resp => {
          //更新前端维护的数据
          let i = this.projects.findIndex(item => {
            return item.url_key == resp.data.data.url_key;
          });
          this.projects[i].name = resp.data.data.name;
          this.projects[i].desc = resp.data.data.desc;
          this.projects[i].updated_at = "刚刚";
          //关闭模态框
          this.is_show_update_dialog = false;
          //提示消息
          this.$message({
            type: "success",
            message: "配置成功"
          });
        })
        .catch(err => {
          console.log(err);
        });
    },
    //创建项目
    createProject() {
      this.$axios({
        method: "POST",
        url: "/api/v1/project",
        data: this.form
      })
        .then(resp => {
          //更新前端维护的数据
          this.projects.push(resp.data.data);
          //关闭模态框
          this.is_show_create_dialog = false;
          //提示消息
          this.$message({
            type: "success",
            message: "创建成功"
          });
        })
        .catch(err => {
          console.log(err);
        });
    },
    //删除项目
    deleteProject(id) {
      //更新后端维护的数据
      this.$axios({
        method: "DELETE",
        url: `/api/v1/project/${id}`
      })
        .then(() => {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          //更新前端维护的数据
          let index = this.projects.findIndex(item => {
            return item.id == id;
          });
          this.projects.splice(index, 1);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>