webpackJsonp([19],{BepB:function(e,t,r){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var s=r("vMJZ"),a={name:"userList",data:function(){var e=this;return{list:null,total:0,listLoading:!0,listQuery:{page:1,number:15,username:void 0,status:void 0,role:void 0},roleOptions:["administrator","writer","subscriber"],statusOptions:[{label:"normal",key:1},{label:"freezing",key:2}],ruleForm:{id:void 0,account:"",nickname:"",email:"",role:"",password:"",passwordAgain:"",website:""},dialogFormVisible:!1,dialogStatus:"",textMap:{update:"edit",create:"create"},rules:{account:[{required:!0,message:this.$t("user.pleaseInputAccount"),trigger:"blur"},{min:3,message:this.$t("user.pleaseCheckAcountLength"),trigger:"blur"}],email:[{required:!0,message:this.$t("user.pleaseInputEmail"),trigger:["blur","change"]},{type:"email",message:this.$t("user.pleaseInputCorrectEmail"),trigger:["blur","change"]}],role:[{required:!0,message:this.$t("user.pleaseSelectRoles"),trigger:"blur"}],password:[{required:!0,validator:function(t,r,s){""===r?s(new Error(e.$t("user.pleaseInputPassWord"))):(""!==e.ruleForm.passwordAgain&&e.$refs.ruleForm.validateField("passwordAgain"),s())},trigger:["blur","change"]},{min:5,message:this.$t("user.pleaseCheckPasswordLength"),trigger:["blur","change"]}],passwordAgain:[{required:!0,validator:function(t,r,s){""===r?s(new Error(e.$t("user.pleaseInputPassWordAgain"))):r!==e.ruleForm.password?s(new Error(e.$t("user.checkPasswordFailed"))):s()},trigger:["blur","change"]},{min:5,message:this.$t("user.pleaseCheckPasswordLength"),trigger:["blur","change"]}]}}},filters:{statusFilter:function(e){return{1:"success",2:"danger"}[e]}},created:function(){this.getList(),this.setTitle()},methods:{setTitle:function(){document.title=this.$t("route."+this.$route.meta.title)+" | Puti"},getList:function(){var e=this;this.listLoading=!0,Object(s.c)(this.listQuery).then(function(t){e.list=t.data.userList,e.total=t.data.totalCount,e.listLoading=!1})},handleFilter:function(){this.listQuery.page=1,this.getList()},handleSizeChange:function(e){this.listQuery.number=e,this.getList()},handleCurrentChange:function(e){this.listQuery.page=e,this.getList()},resetRuleForm:function(){this.ruleForm={id:void 0,account:void 0,nickname:"",email:"",role:"",password:void 0,passwordAgain:void 0,website:""}},handleCreate:function(){var e=this;this.resetRuleForm(),this.dialogStatus="create",this.dialogFormVisible=!0,this.$nextTick(function(){e.$refs.ruleForm.clearValidate()})},createUser:function(){var e=this;this.$refs.ruleForm.validate(function(t){if(!t)return!1;Object(s.a)(e.ruleForm).then(function(t){e.getList(),e.dialogFormVisible=!1,e.$message({message:e.$t("common.createSucceeded"),type:"success",duration:2e3})})})},resetForm:function(e){this.$refs[e].resetFields()},handleUpdate:function(e){var t=this;Object(s.d)(e.account).then(function(e){t.ruleForm.id=e.data.id,t.ruleForm.account=e.data.account,t.ruleForm.nickname=e.data.nickname,t.ruleForm.email=e.data.email,t.ruleForm.role=e.data.roles,t.ruleForm.website=e.data.website,t.ruleForm.password=void 0,t.ruleForm.passwordAgain=void 0}),this.dialogStatus="update",this.dialogFormVisible=!0,this.$nextTick(function(){t.$refs.ruleForm.clearValidate()})},updateUser:function(){var e=this;this.$refs.ruleForm.validate(function(t){t&&Object(s.e)(e.ruleForm).then(function(t){e.dialogFormVisible=!1,0===t.code?(e.getList(),e.$notify({title:e.$t("common.success"),message:e.$t("common.updateSucceeded"),type:"success",duration:2e3})):10002===t.code?e.$notify.error({title:e.$t("common.failed"),message:e.$t("common.updateFailed")+e.$t("common.needRequiredParams"),duration:2e3}):e.$notify.error({title:e.$t("common.failed"),message:e.$t("common.updateFailed")+t.message,duration:2e3})})})},handleModifyStatus:function(e,t){var r,a,i=this;"freeze"===t?(r=2,a={id:e.id,status:r}):(r=1,a={id:e.id,status:r}),Object(s.e)(a).then(function(e){i.$message({message:i.$t("common.operationSucceeded"),type:"success",duration:2e3})}),e.status=r},handleDelete:function(e){var t=this;this.$confirm(this.$t("user.checkToDeleteUser"),this.$t("common.tips"),{confirmButtonText:this.$t("common.confirm"),cancelButtonText:this.$t("common.cancel"),type:"warning",center:!0}).then(function(){Object(s.b)(e.id).then(function(e){t.getList(),t.$message({type:"success",message:t.$t("common.deleteSucceeded")})})}).catch(function(){t.$message({type:"info",message:t.$t("common.cancelDelete")})})}}},i={render:function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("div",{staticClass:"filter-container"},[r("el-input",{staticClass:"filter-item",staticStyle:{width:"200px"},attrs:{placeholder:e.$t("user.username")},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleFilter(t)}},model:{value:e.listQuery.username,callback:function(t){e.$set(e.listQuery,"username",t)},expression:"listQuery.username"}}),e._v(" "),r("el-select",{staticClass:"filter-item",staticStyle:{width:"100px"},attrs:{clearable:"",placeholder:e.$t("user.status")},model:{value:e.listQuery.status,callback:function(t){e.$set(e.listQuery,"status",t)},expression:"listQuery.status"}},e._l(e.statusOptions,function(t){return r("el-option",{key:t.label,attrs:{label:e.$t("user."+t.label),value:t.key}})}),1),e._v(" "),r("el-select",{staticClass:"filter-item",staticStyle:{width:"150px"},attrs:{clearable:""},model:{value:e.listQuery.role,callback:function(t){e.$set(e.listQuery,"role",t)},expression:"listQuery.role"}},e._l(e.roleOptions,function(t){return r("el-option",{key:t,attrs:{label:e.$t("user."+t),value:t}})}),1),e._v(" "),r("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.handleFilter}},[e._v(e._s(e.$t("common.search")))]),e._v(" "),r("el-button",{staticClass:"filter-item",staticStyle:{"margin-left":"10px"},attrs:{type:"primary",icon:"el-icon-edit"},on:{click:e.handleCreate}},[e._v(e._s(e.$t("common.add")))])],1),e._v(" "),r("el-table",{directives:[{name:"loading",rawName:"v-loading.body",value:e.listLoading,expression:"listLoading",modifiers:{body:!0}}],staticStyle:{width:"100%"},attrs:{data:e.list,border:"",fit:"","highlight-current-row":""}},[r("el-table-column",{attrs:{align:"center",label:e.$t("common.ID"),width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.id))])]}}])}),e._v(" "),r("el-table-column",{attrs:{"min-width":"150px",label:e.$t("user.account")},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.account))])]}}])}),e._v(" "),r("el-table-column",{attrs:{"min-width":"150px",label:e.$t("user.nickname")},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.nickname))])]}}])}),e._v(" "),r("el-table-column",{attrs:{"min-width":"150px",label:e.$t("user.email")},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.email))])]}}])}),e._v(" "),r("el-table-column",{attrs:{width:"180px",align:"center",label:e.$t("user.registeredTime")},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(t.row.registered_time))])]}}])}),e._v(" "),r("el-table-column",{attrs:{width:"120px",align:"center",label:e.$t("user.role")},scopedSlots:e._u([{key:"default",fn:function(t){return[r("span",[e._v(e._s(e.$t("user."+t.row.roles)))])]}}])}),e._v(" "),r("el-table-column",{attrs:{align:"center","class-name":"status-col",label:e.$t("user.status"),width:"110"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-tag",{attrs:{type:e._f("statusFilter")(t.row.status)}},[1===t.row.status?r("div",[e._v("\n            "+e._s(e.$t("user.normal"))+"\n          ")]):2===t.row.status?r("div",[e._v("\n            "+e._s(e.$t("user.freezing"))+"\n          ")]):r("div",[e._v("\n             "+e._s(e.$t("common.error"))+"\n          ")])])]}}])}),e._v(" "),r("el-table-column",{attrs:{align:"center",label:e.$t("post.action"),width:"300"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{type:"primary",size:"mini",icon:"el-icon-edit"},on:{click:function(r){return e.handleUpdate(t.row)}}},[e._v(e._s(e.$t("common.edit")))]),e._v(" "),1==t.row.status?r("el-button",{attrs:{size:"mini",type:"warning"},on:{click:function(r){return e.handleModifyStatus(t.row,"freeze")}}},[e._v("\n            "+e._s(e.$t("user.freezing"))+"\n          ")]):2==t.row.status?r("el-button",{attrs:{size:"mini",type:"info"},on:{click:function(r){return e.handleModifyStatus(t.row,"defreeze")}}},[e._v("\n            "+e._s(e.$t("user.defreeze"))+"\n          ")]):e._e(),e._v(" "),r("el-button",{attrs:{type:"danger",size:"mini",icon:"el-icon-delete"},on:{click:function(r){return e.handleDelete(t.row)}}},[e._v(e._s(e.$t("common.delete")))])]}}])})],1),e._v(" "),r("div",{staticClass:"pagination-container"},[r("el-pagination",{attrs:{background:"","current-page":e.listQuery.page,"page-sizes":[10,15,20,30,50],"page-size":e.listQuery.number,layout:"total, sizes, prev, pager, next, jumper",total:e.total},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1),e._v(" "),r("el-dialog",{attrs:{title:e.$t("common."+e.textMap[e.dialogStatus]),visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[r("el-form",{ref:"ruleForm",staticClass:"ruleFrom",attrs:{model:e.ruleForm,rules:e.rules,"label-width":"100px",size:"medium"}},[r("el-form-item",{attrs:{label:e.$t("user.account"),prop:"account"}},["create"==e.dialogStatus?r("el-input",{model:{value:e.ruleForm.account,callback:function(t){e.$set(e.ruleForm,"account",t)},expression:"ruleForm.account"}}):r("el-input",{attrs:{disabled:"disabled"},model:{value:e.ruleForm.account,callback:function(t){e.$set(e.ruleForm,"account",t)},expression:"ruleForm.account"}})],1),e._v(" "),r("el-form-item",{attrs:{label:e.$t("user.nickname"),prop:"nickname"}},[r("el-input",{model:{value:e.ruleForm.nickname,callback:function(t){e.$set(e.ruleForm,"nickname",t)},expression:"ruleForm.nickname"}})],1),e._v(" "),r("el-form-item",{attrs:{label:e.$t("user.email"),prop:"email"}},[r("el-input",{model:{value:e.ruleForm.email,callback:function(t){e.$set(e.ruleForm,"email",t)},expression:"ruleForm.email"}})],1),e._v(" "),r("el-form-item",{attrs:{label:e.$t("user.role"),prop:"role"}},[r("el-select",{attrs:{placeholder:e.$t("user.selectRole")},model:{value:e.ruleForm.role,callback:function(t){e.$set(e.ruleForm,"role",t)},expression:"ruleForm.role"}},[r("el-option",{attrs:{label:e.$t("user.administrator"),value:"administrator"}}),e._v(" "),r("el-option",{attrs:{label:e.$t("user.writer"),value:"writer"}}),e._v(" "),r("el-option",{attrs:{label:e.$t("user.subscriber"),value:"subscriber"}})],1)],1),e._v(" "),"create"==e.dialogStatus?r("el-form-item",{attrs:{label:e.$t("user.password"),prop:"password"}},["create"==e.dialogStatus?r("el-input",{attrs:{type:"password","auto-complete":"off"},model:{value:e.ruleForm.password,callback:function(t){e.$set(e.ruleForm,"password",t)},expression:"ruleForm.password"}}):e._e()],1):e._e(),e._v(" "),"create"==e.dialogStatus?r("el-form-item",{attrs:{label:e.$t("user.passwordAgain"),prop:"passwordAgain"}},[r("el-input",{attrs:{type:"password","auto-complete":"off"},model:{value:e.ruleForm.passwordAgain,callback:function(t){e.$set(e.ruleForm,"passwordAgain",t)},expression:"ruleForm.passwordAgain"}})],1):e._e(),e._v(" "),r("el-form-item",{attrs:{label:e.$t("user.website"),prop:"website"}},[r("el-input",{model:{value:e.ruleForm.website,callback:function(t){e.$set(e.ruleForm,"website",t)},expression:"ruleForm.website"}})],1),e._v(" "),r("el-form-item",["create"==e.dialogStatus?r("el-button",{attrs:{type:"primary"},on:{click:e.createUser}},[e._v(e._s(e.$t("user.createNow")))]):r("el-button",{attrs:{type:"primary"},on:{click:e.updateUser}},[e._v(e._s(e.$t("common.save")))]),e._v(" "),"create"==e.dialogStatus?r("el-button",{on:{click:function(t){return e.resetForm("ruleForm")}}},[e._v(e._s(e.$t("common.reset")))]):r("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v(e._s(e.$t("common.cancel")))])],1)],1)],1)],1)},staticRenderFns:[]},l=r("VU/8")(a,i,!1,null,null,null);t.default=l.exports},vMJZ:function(e,t,r){"use strict";t.c=function(e){return Object(s.a)({url:"/user",method:"get",params:e})},t.d=function(e){return Object(s.a)({url:"/user/"+e,method:"get"})},t.a=function(e){return Object(s.a)({url:"/user/"+e.account,method:"post",data:e})},t.e=function(e){return Object(s.a)({url:"/user/"+e.id,method:"put",data:e})},t.b=function(e){return Object(s.a)({url:"/user/"+e,method:"delete"})};var s=r("vLgD")}});