(this.webpackJsonpweb=this.webpackJsonpweb||[]).push([[0],{10:function(e,t,n){},13:function(e,t,n){"use strict";n.r(t);var c=n(1),s=n.n(c),r=n(4),a=n.n(r),i=(n(9),n(2)),o=(n(10),n(11),n(0));var u=function(){var e=Object(c.useState)(""),t=Object(i.a)(e,2),n=t[0],s=t[1],r=Object(c.useState)([]),a=Object(i.a)(r,2),u=a[0],p=a[1],j=Object(c.useState)({open:!1,msg:""}),l=Object(i.a)(j,2),d=l[0],m=l[1],b=Object(c.useState)(!1),h=Object(i.a)(b,2),O=h[0],g=h[1];return Object(o.jsxs)("div",{className:"App",children:[Object(o.jsx)("div",{className:"App-header",children:Object(o.jsxs)("div",{children:[Object(o.jsx)("input",{placeholder:"mdpr url",value:n,onChange:function(e){return function(e){s(e.target.value)}(e)},className:"App-input"}),Object(o.jsx)("button",{className:"primary",onClick:function(){return function(){if(m({open:!1,msg:""}),g(!0),!/http(s)?:\/\/([\w-]+.)+[\w-]+(\/[\w- ./?%&=]*)?/.test(n)||-1===n.indexOf("mdpr.jp"))return g(!1),m({open:!0,msg:"URL Error"}),!1;var e=n.split(" "),t=e[e.length-1];fetch("/api/mdpr?url="+t,{method:"GET",dataType:"json"}).then((function(e){return e.json()})).then((function(e){1===e.status?(g(!1),m({open:!1,msg:""}),p(e.data)):(g(!1),m({open:!0,msg:e.message}))})).catch((function(){g(!1),m({open:!0,msg:"Server Error"})}))}()},children:"GET"})]})}),Object(o.jsx)("div",{className:"App-main",children:O?Object(o.jsx)("div",{className:"spinner primary"}):d.open?Object(o.jsx)("button",{className:"secondary",children:d.msg}):u.map((function(e,t){return Object(o.jsx)("div",{children:Object(o.jsx)("img",{src:e,alt:"",className:"App-result-img"})},t)}))})]})},p=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,14)).then((function(t){var n=t.getCLS,c=t.getFID,s=t.getFCP,r=t.getLCP,a=t.getTTFB;n(e),c(e),s(e),r(e),a(e)}))};a.a.render(Object(o.jsx)(s.a.StrictMode,{children:Object(o.jsx)(u,{})}),document.getElementById("root")),p()},9:function(e,t,n){}},[[13,1,2]]]);
//# sourceMappingURL=main.3aa10aad.chunk.js.map