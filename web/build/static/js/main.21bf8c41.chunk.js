(this.webpackJsonpweb=this.webpackJsonpweb||[]).push([[0],{10:function(e,t,n){},12:function(e,t,n){"use strict";n.r(t);var c=n(1),r=n.n(c),a=n(4),s=n.n(a),i=(n(9),n(3)),u=(n(10),n(0));var o=function(){var e=Object(c.useState)(""),t=Object(i.a)(e,2),n=t[0],r=t[1],a=Object(c.useState)([]),s=Object(i.a)(a,2),o=s[0],l=s[1];return Object(u.jsxs)("div",{className:"App",children:[Object(u.jsx)("header",{className:"App-header",children:Object(u.jsxs)("div",{children:[Object(u.jsx)("input",{placeholder:"mdpr url",value:n,onChange:function(e){return function(e){r(e.target.value)}(e)},className:"App-input"}),Object(u.jsx)("button",{onClick:function(){return function(){if(!/http(s)?:\/\/([\w-]+.)+[\w-]+(\/[\w- ./?%&=]*)?/.test(n)||-1===n.indexOf("mdpr.jp"))return alert("URL Error"),!1;var e=n.split(" "),t=e[e.length-1];fetch("/api/mdpr?url="+t,{method:"GET",dataType:"json"}).then((function(e){return e.json()})).then((function(e){1===e.status?l(e.data):alert(e.message)})).catch((function(){alert("Server Error")}))}()},children:"GET"})]})}),Object(u.jsx)("main",{className:"App-main",children:o.map((function(e,t){return Object(u.jsx)("div",{children:Object(u.jsx)("img",{src:e,alt:"",className:"App-result-img"})},t)}))})]})},l=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,13)).then((function(t){var n=t.getCLS,c=t.getFID,r=t.getFCP,a=t.getLCP,s=t.getTTFB;n(e),c(e),r(e),a(e),s(e)}))};s.a.render(Object(u.jsx)(r.a.StrictMode,{children:Object(u.jsx)(o,{})}),document.getElementById("root")),l()},9:function(e,t,n){}},[[12,1,2]]]);
//# sourceMappingURL=main.21bf8c41.chunk.js.map