(this.webpackJsonpweb=this.webpackJsonpweb||[]).push([[0],{10:function(t,e,n){},12:function(t,e,n){"use strict";n.r(e);var c=n(1),r=n.n(c),a=n(4),i=n.n(a),s=(n(9),n(3)),u=(n(10),n(0));var o=function(){var t=Object(c.useState)(""),e=Object(s.a)(t,2),n=e[0],r=e[1],a=Object(c.useState)([]),i=Object(s.a)(a,2),o=i[0],j=i[1];return Object(u.jsxs)("div",{className:"App",children:[Object(u.jsxs)("header",{className:"App-header",children:[Object(u.jsx)("div",{children:Object(u.jsx)("input",{value:n,onChange:function(t){return function(t){r(t.target.value)}(t)}})}),Object(u.jsx)("div",{children:Object(u.jsx)("button",{onClick:function(){return function(){if(!/http(s)?:\/\/([\w-]+.)+[\w-]+(\/[\w- ./?%&=]*)?/.test(n)||-1===n.indexOf("mdpr.jp"))return alert("URL Error"),!1;var t=n.split(" "),e=t[t.length-1];fetch("/api/mdpr?url="+e,{method:"GET",dataType:"json"}).then((function(t){return t.json()})).then((function(t){1===t.status?j(t.data):alert(t.message)})).catch((function(){alert("Server Error")}))}()},children:"Get"})})]}),Object(u.jsx)("main",{className:"App-main",children:o.map((function(t,e){return Object(u.jsx)("div",{children:Object(u.jsx)("img",{src:t,alt:"",className:"App-result-img"})},e)}))})]})},j=function(t){t&&t instanceof Function&&n.e(3).then(n.bind(null,13)).then((function(e){var n=e.getCLS,c=e.getFID,r=e.getFCP,a=e.getLCP,i=e.getTTFB;n(t),c(t),r(t),a(t),i(t)}))};i.a.render(Object(u.jsx)(r.a.StrictMode,{children:Object(u.jsx)(o,{})}),document.getElementById("root")),j()},9:function(t,e,n){}},[[12,1,2]]]);
//# sourceMappingURL=main.7977db97.chunk.js.map