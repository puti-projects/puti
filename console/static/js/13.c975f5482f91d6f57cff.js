webpackJsonp([13],{Exc4:function(t,a,e){var n=e("m22A");"string"==typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);e("rjj0")("fa598d04",n,!0)},eRLo:function(t,a,e){"use strict";Object.defineProperty(a,"__esModule",{value:!0});var n=e("mw02"),i=e.n(n),r={name:"page401",data:function(){return{errGif:i.a+"?"+ +new Date,ewizardClap:"https://wpimg.wallstcn.com/007ef517-bafd-4066-aae4-6883632d9646",dialogVisible:!1}},methods:{back:function(){this.$route.query.noGoBack?this.$router.push({path:"/dashboard"}):this.$router.go(-1)}}},o={render:function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"errPage-container"},[e("el-button",{staticClass:"pan-back-btn",attrs:{icon:"arrow-left"},on:{click:t.back}},[t._v("返回")]),t._v(" "),e("el-row",[e("el-col",{attrs:{span:12}},[e("h1",{staticClass:"text-jumbo text-ginormous"},[t._v("Oops!")]),t._v("\n      gif来源"),e("a",{attrs:{href:"https://zh.airbnb.com/",target:"_blank"}},[t._v("airbnb")]),t._v(" 页面\n      "),e("h2",[t._v("你没有权限去该页面")]),t._v(" "),e("h6",[t._v("如有不满请联系你领导")]),t._v(" "),e("ul",{staticClass:"list-unstyled"},[e("li",[t._v("或者你可以去:")]),t._v(" "),e("li",{staticClass:"link-type"},[e("router-link",{attrs:{to:"/dashboard"}},[t._v("回首页")])],1),t._v(" "),e("li",{staticClass:"link-type"},[e("a",{attrs:{href:"https://www.taobao.com/"}},[t._v("随便看看")])]),t._v(" "),e("li",[e("a",{attrs:{href:"#"},on:{click:function(a){a.preventDefault(),t.dialogVisible=!0}}},[t._v("点我看图")])])])]),t._v(" "),e("el-col",{attrs:{span:12}},[e("img",{attrs:{src:t.errGif,width:"313",height:"428",alt:"Girl has dropped her ice cream."}})])],1),t._v(" "),e("el-dialog",{attrs:{title:"随便看",visible:t.dialogVisible},on:{"update:visible":function(a){t.dialogVisible=a}}},[e("img",{staticClass:"pan-img",attrs:{src:t.ewizardClap}})])],1)},staticRenderFns:[]};var s=e("VU/8")(r,o,!1,function(t){e("Exc4")},"data-v-1ee28410",null);a.default=s.exports},m22A:function(t,a,e){(t.exports=e("FZ+f")(!1)).push([t.i,"\n.errPage-container[data-v-1ee28410] {\n  width: 800px;\n  margin: 100px auto;\n}\n.errPage-container .pan-back-btn[data-v-1ee28410] {\n    background: #008489;\n    color: #fff;\n    border: none !important;\n}\n.errPage-container .pan-gif[data-v-1ee28410] {\n    margin: 0 auto;\n    display: block;\n}\n.errPage-container .pan-img[data-v-1ee28410] {\n    display: block;\n    margin: 0 auto;\n    width: 100%;\n}\n.errPage-container .text-jumbo[data-v-1ee28410] {\n    font-size: 60px;\n    font-weight: 700;\n    color: #484848;\n}\n.errPage-container .list-unstyled[data-v-1ee28410] {\n    font-size: 14px;\n}\n.errPage-container .list-unstyled li[data-v-1ee28410] {\n      padding-bottom: 5px;\n}\n.errPage-container .list-unstyled a[data-v-1ee28410] {\n      color: #008489;\n      text-decoration: none;\n}\n.errPage-container .list-unstyled a[data-v-1ee28410]:hover {\n        text-decoration: underline;\n}\n",""])},mw02:function(t,a,e){t.exports=e.p+"static/img/401.089007e.gif"}});