webpackJsonp([7],{"22fj":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=n("/5sW"),o=n("TQvf"),r=n.n(o);function a(t,e){var n=new r.a(e.target,{text:function(){return t}});n.on("success",function(){i.default.prototype.$message({message:"Copy successfully",type:"success",duration:1500}),n.off("error"),n.off("success"),n.destroy()}),n.on("error",function(){i.default.prototype.$message({message:"Copy failed",type:"error"}),n.off("error"),n.off("success"),n.destroy()}),n.onClick(e)}var l=n("oCAU"),s={name:"mediaLiarary",data:function(){return{list:null,total:0,listLoading:!0,listQuery:{page:1,limit:24},dialogVisible:!1,dialogTitle:"File",detailForm:{id:void 0,title:"",slug:"",url:"",type:"",uploadTime:"",description:""}}},created:function(){this.getList(),this.setTitle()},methods:{setTitle:function(){document.title=this.$t("route."+this.$route.meta.title)+" | Puti"},getList:function(){var t=this;this.listLoading=!0,Object(l.b)(this.listQuery).then(function(e){t.list=e.data.mediaList,t.total=e.data.totalCount,t.listLoading=!1})},handleSizeChange:function(t){this.listQuery.limit=t,this.getList()},handleCurrentChange:function(t){this.listQuery.page=t,this.getList()},handleDelete:function(t){var e=this;this.$confirm(this.$t("media.checkToDeleteMedia")+this.$t("media.fileName")+":"+t.title,this.$t("common.tips"),{confirmButtonText:this.$t("common.confirm"),cancelButtonText:this.$t("common.cancel"),type:"warning",center:!0}).then(function(){Object(l.a)(t.id).then(function(n){if(0===n.code){e.$message({message:e.$t("common.deleteSucceeded"),type:"success",duration:3e3});var i=e.list.indexOf(t);e.list.splice(i,1)}else e.$message.error({message:e.$t("common.operationFailed")+n.message,duration:3e3})})}).catch(function(){e.$message({type:"info",message:e.$t("common.cancelDelete")})})},handleDetail:function(t){var e=this;Object(l.c)(t.id).then(function(n){0===n.code?(e.detailForm.id=n.data.id,e.detailForm.title=n.data.title,e.detailForm.slug=n.data.slug,e.detailForm.url=n.data.url,e.detailForm.type=n.data.type,e.detailForm.uploadTime=n.data.upload_time,e.detailForm.description=n.data.description,e.dialogVisible=!0,e.dialogTitle=t.title,e.dialogImgPreviewUrl=t.url):e.$message.error({message:e.$t("media.getDetailFailed")+n.message,duration:3e3})})},handleCopy:function(t,e){a(t,e)},clipboardSuccess:function(){this.$message({message:this.$t("media.copySuccessed"),type:"success",duration:1500})},handleUpdate:function(){var t=this,e={title:this.detailForm.title,slug:this.detailForm.slug,description:this.detailForm.description};Object(l.d)(this.detailForm.id,e).then(function(e){0===e.code?(t.$message({message:t.$t("common.updateSucceeded"),type:"success",duration:3e3}),t.dialogVisible=!1,t.getList()):20203===e.code?t.$message.error({message:t.$t("media.titleEmpty"),duration:3e3}):t.$message.error({message:t.$t("common.updateFailed")+e.message,duration:3e3})})}}},c={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("el-row",{directives:[{name:"loading",rawName:"v-loading.body",value:t.listLoading,expression:"listLoading",modifiers:{body:!0}}]},t._l(t.list,function(e){return n("el-col",{key:e.id,attrs:{xs:24,sm:8,md:4,xl:3}},[n("el-card",{staticClass:"media-card",attrs:{"body-style":{padding:"0px"}}},["picture"===e.type?n("img",{staticClass:"media-image",attrs:{src:e.url}}):n("div",{staticClass:"media-other-file",attrs:{icon:"el-icon-document"}},[n("svg-icon",{staticClass:"media-other-file-svg",attrs:{"icon-class":"article"}})],1),t._v(" "),n("div",{staticStyle:{padding:"8px",width:"100%"}},[n("el-button",{staticClass:"media-title",attrs:{type:"text"},on:{click:function(n){t.handleDetail(e)}}},[t._v(t._s(e.title))]),t._v(" "),n("div",{staticClass:"media-bottom clearfix"},[n("time",{staticClass:"media-time"},[t._v(t._s(e.upload_time))]),t._v(" "),n("el-button",{staticClass:"media-button",on:{click:function(n){t.handleDelete(e)}}},[n("i",{staticClass:"el-icon-delete"})])],1)],1)])],1)})),t._v(" "),n("div",{staticClass:"pagination-container"},[n("el-pagination",{attrs:{background:"","current-page":t.listQuery.page,"page-sizes":[16,24,32,40],"page-size":t.listQuery.limit,layout:"total, sizes, prev, pager, next, jumper",total:t.total},on:{"size-change":t.handleSizeChange,"current-change":t.handleCurrentChange}})],1),t._v(" "),n("el-dialog",{attrs:{title:t.dialogTitle,visible:t.dialogVisible,width:"60%"},on:{"update:visible":function(e){t.dialogVisible=e}}},[n("el-row",{attrs:{gutter:20}},[n("el-col",{staticClass:"media-detail-left",attrs:{xs:24,sm:24,md:15,xl:15}},["picture"==t.detailForm.type?n("div",{staticClass:"grid-content"},[n("img",{staticClass:"media-detail-image",attrs:{src:t.detailForm.url}})]):n("div",{staticClass:"grid-content"},[t._v(t._s(t.$t("media.unknowType")))])]),t._v(" "),n("el-col",{attrs:{xs:24,sm:24,md:9,xl:9}},[n("div",{staticClass:"grid-content"},[n("el-form",{attrs:{model:t.detailForm}},[n("el-form-item",{attrs:{label:t.$t("media.mediaTitle")}},[n("el-input",{model:{value:t.detailForm.title,callback:function(e){t.$set(t.detailForm,"title",e)},expression:"detailForm.title"}})],1),t._v(" "),n("el-form-item",{attrs:{label:t.$t("media.mediaSlug")}},[n("el-input",{model:{value:t.detailForm.slug,callback:function(e){t.$set(t.detailForm,"slug",e)},expression:"detailForm.slug"}})],1),t._v(" "),n("el-form-item",{attrs:{label:t.$t("media.mediaUrl")}},[n("el-tooltip",{attrs:{effect:"dark",content:t.$t("media.copyTo"),placement:"right-end"}},[n("el-button",{staticClass:"media-url-copy el-icon-document",attrs:{type:"text",size:"mini"},on:{click:function(e){t.handleCopy(t.detailForm.url,e)}}})],1),t._v(" "),n("el-input",{attrs:{disabled:"disabled"},model:{value:t.detailForm.url,callback:function(e){t.$set(t.detailForm,"url",e)},expression:"detailForm.url"}})],1),t._v(" "),n("el-form-item",{attrs:{label:t.$t("media.mediaDescription")}},[n("el-input",{attrs:{type:"textarea",autosize:{minRows:2,maxRows:4},placeholder:t.$t("media.pleaseInputDesc")},model:{value:t.detailForm.description,callback:function(e){t.$set(t.detailForm,"description",e)},expression:"detailForm.description"}})],1),t._v(" "),n("el-form-item",{attrs:{label:t.$t("media.mediaUploadTime")}},[n("span",[t._v(t._s(t.detailForm.uploadTime))])]),t._v(" "),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:function(e){t.handleUpdate()}}},[t._v(t._s(t.$t("common.change")))])],1)],1)],1)])],1)],1)],1)},staticRenderFns:[]};var u=n("VU/8")(s,c,!1,function(t){n("Md8S")},null,null);e.default=u.exports},Md8S:function(t,e,n){var i=n("P0IZ");"string"==typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);n("rjj0")("0c4d26d3",i,!0)},P0IZ:function(t,e,n){(t.exports=n("FZ+f")(!1)).push([t.i,'\n.media-card{\n  margin: 10px;\n}\n.media-title{\n  display: inline-block;\n  overflow: hidden;\n  text-overflow:ellipsis;\n  white-space: nowrap;\n  width:100%;\n  padding: 0;\n  text-align: left;\n}\n.media-time {\n  font-size: 13px;\n  color: #999;\n}\n.media-bottom {\n  margin-top: 13px;\n  line-height: 12px;\n}\n.media-button {\n  padding: 0;\n  float: right;\n  border: 0;\n  color: #F56C6C;\n}\n.media-button:hover {\n  color: rgb(255, 0, 0);\n  background-color: #fff;\n}\n.media-button:focus{\n  color: #fff;\n  background-color: rgb(255, 0, 0);\n}\n.media-image {\n  width: 100%;\n  height: 180px;\n  display: block;\n}\n.media-other-file{\n  width: 100%;\n  height: 180px;\n  display: block;\n}\n.media-other-file-svg{\n  display: block;\n  font-size: 140px;\n  margin: 0 auto;\n  padding-top: 40px;\n  color: #5B584A;\n  overflow: hidden;\n}\n.media-detail-image{\n  margin: 0 auto;\n  width: 100%;\n  height: 100%;\n  display: -webkit-box;\n  display: -ms-flexbox;\n  display: flex;\n  -webkit-box-align: center;\n      -ms-flex-align: center;\n          align-items: center;\n}\n.media-url-copy{\n  color: #999\n}\n.clearfix:before,\n.clearfix:after {\n    display: table;\n    content: "";\n}\n.clearfix:after {\n    clear: both\n}\n',""])},TQvf:function(t,e,n){var i;i=function(){return function(t){var e={};function n(i){if(e[i])return e[i].exports;var o=e[i]={i:i,l:!1,exports:{}};return t[i].call(o.exports,o,o.exports,n),o.l=!0,o.exports}return n.m=t,n.c=e,n.i=function(t){return t},n.d=function(t,e,i){n.o(t,e)||Object.defineProperty(t,e,{configurable:!1,enumerable:!0,get:i})},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="",n(n.s=3)}([function(t,e,n){var i,o,r,a;a=function(t,e){"use strict";var n,i=(n=e)&&n.__esModule?n:{default:n};var o="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t};var r=function(){function t(t,e){for(var n=0;n<e.length;n++){var i=e[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(t,i.key,i)}}return function(e,n,i){return n&&t(e.prototype,n),i&&t(e,i),e}}(),a=function(){function t(e){!function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}(this,t),this.resolveOptions(e),this.initSelection()}return r(t,[{key:"resolveOptions",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};this.action=t.action,this.container=t.container,this.emitter=t.emitter,this.target=t.target,this.text=t.text,this.trigger=t.trigger,this.selectedText=""}},{key:"initSelection",value:function(){this.text?this.selectFake():this.target&&this.selectTarget()}},{key:"selectFake",value:function(){var t=this,e="rtl"==document.documentElement.getAttribute("dir");this.removeFake(),this.fakeHandlerCallback=function(){return t.removeFake()},this.fakeHandler=this.container.addEventListener("click",this.fakeHandlerCallback)||!0,this.fakeElem=document.createElement("textarea"),this.fakeElem.style.fontSize="12pt",this.fakeElem.style.border="0",this.fakeElem.style.padding="0",this.fakeElem.style.margin="0",this.fakeElem.style.position="absolute",this.fakeElem.style[e?"right":"left"]="-9999px";var n=window.pageYOffset||document.documentElement.scrollTop;this.fakeElem.style.top=n+"px",this.fakeElem.setAttribute("readonly",""),this.fakeElem.value=this.text,this.container.appendChild(this.fakeElem),this.selectedText=(0,i.default)(this.fakeElem),this.copyText()}},{key:"removeFake",value:function(){this.fakeHandler&&(this.container.removeEventListener("click",this.fakeHandlerCallback),this.fakeHandler=null,this.fakeHandlerCallback=null),this.fakeElem&&(this.container.removeChild(this.fakeElem),this.fakeElem=null)}},{key:"selectTarget",value:function(){this.selectedText=(0,i.default)(this.target),this.copyText()}},{key:"copyText",value:function(){var t=void 0;try{t=document.execCommand(this.action)}catch(e){t=!1}this.handleResult(t)}},{key:"handleResult",value:function(t){this.emitter.emit(t?"success":"error",{action:this.action,text:this.selectedText,trigger:this.trigger,clearSelection:this.clearSelection.bind(this)})}},{key:"clearSelection",value:function(){this.trigger&&this.trigger.focus(),window.getSelection().removeAllRanges()}},{key:"destroy",value:function(){this.removeFake()}},{key:"action",set:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"copy";if(this._action=t,"copy"!==this._action&&"cut"!==this._action)throw new Error('Invalid "action" value, use either "copy" or "cut"')},get:function(){return this._action}},{key:"target",set:function(t){if(void 0!==t){if(!t||"object"!==(void 0===t?"undefined":o(t))||1!==t.nodeType)throw new Error('Invalid "target" value, use a valid Element');if("copy"===this.action&&t.hasAttribute("disabled"))throw new Error('Invalid "target" attribute. Please use "readonly" instead of "disabled" attribute');if("cut"===this.action&&(t.hasAttribute("readonly")||t.hasAttribute("disabled")))throw new Error('Invalid "target" attribute. You can\'t cut text from elements with "readonly" or "disabled" attributes');this._target=t}},get:function(){return this._target}}]),t}();t.exports=a},o=[t,n(7)],void 0===(r="function"==typeof(i=a)?i.apply(e,o):i)||(t.exports=r)},function(t,e,n){var i=n(6),o=n(5);t.exports=function(t,e,n){if(!t&&!e&&!n)throw new Error("Missing required arguments");if(!i.string(e))throw new TypeError("Second argument must be a String");if(!i.fn(n))throw new TypeError("Third argument must be a Function");if(i.node(t))return function(t,e,n){return t.addEventListener(e,n),{destroy:function(){t.removeEventListener(e,n)}}}(t,e,n);if(i.nodeList(t))return function(t,e,n){return Array.prototype.forEach.call(t,function(t){t.addEventListener(e,n)}),{destroy:function(){Array.prototype.forEach.call(t,function(t){t.removeEventListener(e,n)})}}}(t,e,n);if(i.string(t))return function(t,e,n){return o(document.body,t,e,n)}(t,e,n);throw new TypeError("First argument must be a String, HTMLElement, HTMLCollection, or NodeList")}},function(t,e){function n(){}n.prototype={on:function(t,e,n){var i=this.e||(this.e={});return(i[t]||(i[t]=[])).push({fn:e,ctx:n}),this},once:function(t,e,n){var i=this;function o(){i.off(t,o),e.apply(n,arguments)}return o._=e,this.on(t,o,n)},emit:function(t){for(var e=[].slice.call(arguments,1),n=((this.e||(this.e={}))[t]||[]).slice(),i=0,o=n.length;i<o;i++)n[i].fn.apply(n[i].ctx,e);return this},off:function(t,e){var n=this.e||(this.e={}),i=n[t],o=[];if(i&&e)for(var r=0,a=i.length;r<a;r++)i[r].fn!==e&&i[r].fn._!==e&&o.push(i[r]);return o.length?n[t]=o:delete n[t],this}},t.exports=n},function(t,e,n){var i,o,r,a;a=function(t,e,n,i){"use strict";var o=l(e),r=l(n),a=l(i);function l(t){return t&&t.__esModule?t:{default:t}}var s="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t};var c=function(){function t(t,e){for(var n=0;n<e.length;n++){var i=e[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(t,i.key,i)}}return function(e,n,i){return n&&t(e.prototype,n),i&&t(e,i),e}}();var u=function(t){function e(t,n){!function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}(this,e);var i=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!=typeof e&&"function"!=typeof e?t:e}(this,(e.__proto__||Object.getPrototypeOf(e)).call(this));return i.resolveOptions(n),i.listenClick(t),i}return function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+typeof e);t.prototype=Object.create(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(Object.setPrototypeOf?Object.setPrototypeOf(t,e):t.__proto__=e)}(e,r.default),c(e,[{key:"resolveOptions",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};this.action="function"==typeof t.action?t.action:this.defaultAction,this.target="function"==typeof t.target?t.target:this.defaultTarget,this.text="function"==typeof t.text?t.text:this.defaultText,this.container="object"===s(t.container)?t.container:document.body}},{key:"listenClick",value:function(t){var e=this;this.listener=(0,a.default)(t,"click",function(t){return e.onClick(t)})}},{key:"onClick",value:function(t){var e=t.delegateTarget||t.currentTarget;this.clipboardAction&&(this.clipboardAction=null),this.clipboardAction=new o.default({action:this.action(e),target:this.target(e),text:this.text(e),container:this.container,trigger:e,emitter:this})}},{key:"defaultAction",value:function(t){return d("action",t)}},{key:"defaultTarget",value:function(t){var e=d("target",t);if(e)return document.querySelector(e)}},{key:"defaultText",value:function(t){return d("text",t)}},{key:"destroy",value:function(){this.listener.destroy(),this.clipboardAction&&(this.clipboardAction.destroy(),this.clipboardAction=null)}}],[{key:"isSupported",value:function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:["copy","cut"],e="string"==typeof t?[t]:t,n=!!document.queryCommandSupported;return e.forEach(function(t){n=n&&!!document.queryCommandSupported(t)}),n}}]),e}();function d(t,e){var n="data-clipboard-"+t;if(e.hasAttribute(n))return e.getAttribute(n)}t.exports=u},o=[t,n(0),n(2),n(1)],void 0===(r="function"==typeof(i=a)?i.apply(e,o):i)||(t.exports=r)},function(t,e){var n=9;if("undefined"!=typeof Element&&!Element.prototype.matches){var i=Element.prototype;i.matches=i.matchesSelector||i.mozMatchesSelector||i.msMatchesSelector||i.oMatchesSelector||i.webkitMatchesSelector}t.exports=function(t,e){for(;t&&t.nodeType!==n;){if("function"==typeof t.matches&&t.matches(e))return t;t=t.parentNode}}},function(t,e,n){var i=n(4);function o(t,e,n,o,r){var a=function(t,e,n,o){return function(n){n.delegateTarget=i(n.target,e),n.delegateTarget&&o.call(t,n)}}.apply(this,arguments);return t.addEventListener(n,a,r),{destroy:function(){t.removeEventListener(n,a,r)}}}t.exports=function(t,e,n,i,r){return"function"==typeof t.addEventListener?o.apply(null,arguments):"function"==typeof n?o.bind(null,document).apply(null,arguments):("string"==typeof t&&(t=document.querySelectorAll(t)),Array.prototype.map.call(t,function(t){return o(t,e,n,i,r)}))}},function(t,e){e.node=function(t){return void 0!==t&&t instanceof HTMLElement&&1===t.nodeType},e.nodeList=function(t){var n=Object.prototype.toString.call(t);return void 0!==t&&("[object NodeList]"===n||"[object HTMLCollection]"===n)&&"length"in t&&(0===t.length||e.node(t[0]))},e.string=function(t){return"string"==typeof t||t instanceof String},e.fn=function(t){return"[object Function]"===Object.prototype.toString.call(t)}},function(t,e){t.exports=function(t){var e;if("SELECT"===t.nodeName)t.focus(),e=t.value;else if("INPUT"===t.nodeName||"TEXTAREA"===t.nodeName){var n=t.hasAttribute("readonly");n||t.setAttribute("readonly",""),t.select(),t.setSelectionRange(0,t.value.length),n||t.removeAttribute("readonly"),e=t.value}else{t.hasAttribute("contenteditable")&&t.focus();var i=window.getSelection(),o=document.createRange();o.selectNodeContents(t),i.removeAllRanges(),i.addRange(o),e=i.toString()}return e}}])},t.exports=i()},oCAU:function(t,e,n){"use strict";e.b=function(t){return Object(i.a)({url:"/media",method:"get",params:t})},e.c=function(t){return Object(i.a)({url:"/media/"+t,method:"get"})},e.a=function(t){return Object(i.a)({url:"/media/"+t,method:"delete"})},e.d=function(t,e){return Object(i.a)({url:"/media/"+t,method:"put",data:e})};var i=n("vLgD")}});