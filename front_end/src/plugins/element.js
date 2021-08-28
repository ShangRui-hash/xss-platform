import Vue from 'vue'
import {
    Alert,
    Tag,
    Button,
    Menu,
    Submenu,
    MenuItem,
    MenuItemGroup,
    Dialog,
    Form,
    FormItem,
    Input,
    Table,
    TableColumn,
    Container,
    Aside,
    Main,
    Dropdown,
    DropdownMenu,
    DropdownItem,
    Footer,
    Message,
    Popconfirm,
    Checkbox,
    CheckboxGroup,
    Card,
    MessageBox,
    Drawer,
    Link,
    Notification,
    Switch
} from 'element-ui';
import '../../theme/index.css'
Vue.use(Alert)
Vue.use(Tag)
Vue.use(Button)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)
Vue.use(Dialog)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Container)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Dropdown)
Vue.use(DropdownMenu)
Vue.use(DropdownItem)
Vue.use(Footer)
Vue.use(Popconfirm)
Vue.use(Checkbox)
Vue.use(CheckboxGroup)
Vue.use(Card)
Vue.use(Drawer)
Vue.use(Link)
Vue.use(Switch)
Vue.prototype.$notify = Notification
Vue.prototype.$message = Message
Vue.prototype.$alert = MessageBox