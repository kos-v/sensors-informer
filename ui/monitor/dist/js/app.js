(function(s){function e(e){for(var t,c,i=e[0],l=e[1],o=e[2],u=0,f=[];u<i.length;u++)c=i[u],Object.prototype.hasOwnProperty.call(n,c)&&n[c]&&f.push(n[c][0]),n[c]=0;for(t in l)Object.prototype.hasOwnProperty.call(l,t)&&(s[t]=l[t]);d&&d(e);while(f.length)f.shift()();return r.push.apply(r,o||[]),a()}function a(){for(var s,e=0;e<r.length;e++){for(var a=r[e],t=!0,i=1;i<a.length;i++){var l=a[i];0!==n[l]&&(t=!1)}t&&(r.splice(e--,1),s=c(c.s=a[0]))}return s}var t={},n={app:0},r=[];function c(e){if(t[e])return t[e].exports;var a=t[e]={i:e,l:!1,exports:{}};return s[e].call(a.exports,a,a.exports,c),a.l=!0,a.exports}c.m=s,c.c=t,c.d=function(s,e,a){c.o(s,e)||Object.defineProperty(s,e,{enumerable:!0,get:a})},c.r=function(s){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(s,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(s,"__esModule",{value:!0})},c.t=function(s,e){if(1&e&&(s=c(s)),8&e)return s;if(4&e&&"object"===typeof s&&s&&s.__esModule)return s;var a=Object.create(null);if(c.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:s}),2&e&&"string"!=typeof s)for(var t in s)c.d(a,t,function(e){return s[e]}.bind(null,t));return a},c.n=function(s){var e=s&&s.__esModule?function(){return s["default"]}:function(){return s};return c.d(e,"a",e),e},c.o=function(s,e){return Object.prototype.hasOwnProperty.call(s,e)},c.p="/dist/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],l=i.push.bind(i);i.push=e,i=i.slice();for(var o=0;o<i.length;o++)e(i[o]);var d=l;r.push([0,"chunk-vendors"]),a()})({0:function(s,e,a){s.exports=a("56d7")},"034f":function(s,e,a){"use strict";a("85ec")},4678:function(s,e,a){var t={"./af":"2bfb","./af.js":"2bfb","./ar":"8e73","./ar-dz":"a356","./ar-dz.js":"a356","./ar-kw":"423e","./ar-kw.js":"423e","./ar-ly":"1cfd","./ar-ly.js":"1cfd","./ar-ma":"0a84","./ar-ma.js":"0a84","./ar-sa":"8230","./ar-sa.js":"8230","./ar-tn":"6d83","./ar-tn.js":"6d83","./ar.js":"8e73","./az":"485c","./az.js":"485c","./be":"1fc1","./be.js":"1fc1","./bg":"84aa","./bg.js":"84aa","./bm":"a7fa","./bm.js":"a7fa","./bn":"9043","./bn-bd":"9686","./bn-bd.js":"9686","./bn.js":"9043","./bo":"d26a","./bo.js":"d26a","./br":"6887","./br.js":"6887","./bs":"2554","./bs.js":"2554","./ca":"d716","./ca.js":"d716","./cs":"3c0d","./cs.js":"3c0d","./cv":"03ec","./cv.js":"03ec","./cy":"9797","./cy.js":"9797","./da":"0f14","./da.js":"0f14","./de":"b469","./de-at":"b3eb","./de-at.js":"b3eb","./de-ch":"bb71","./de-ch.js":"bb71","./de.js":"b469","./dv":"598a","./dv.js":"598a","./el":"8d47","./el.js":"8d47","./en-au":"0e6b","./en-au.js":"0e6b","./en-ca":"3886","./en-ca.js":"3886","./en-gb":"39a6","./en-gb.js":"39a6","./en-ie":"e1d3","./en-ie.js":"e1d3","./en-il":"7333","./en-il.js":"7333","./en-in":"ec2e","./en-in.js":"ec2e","./en-nz":"6f50","./en-nz.js":"6f50","./en-sg":"b7e9","./en-sg.js":"b7e9","./eo":"65db","./eo.js":"65db","./es":"898b","./es-do":"0a3c","./es-do.js":"0a3c","./es-mx":"b5b7","./es-mx.js":"b5b7","./es-us":"55c9","./es-us.js":"55c9","./es.js":"898b","./et":"ec18","./et.js":"ec18","./eu":"0ff2","./eu.js":"0ff2","./fa":"8df4","./fa.js":"8df4","./fi":"81e9","./fi.js":"81e9","./fil":"d69a","./fil.js":"d69a","./fo":"0721","./fo.js":"0721","./fr":"9f26","./fr-ca":"d9f8","./fr-ca.js":"d9f8","./fr-ch":"0e49","./fr-ch.js":"0e49","./fr.js":"9f26","./fy":"7118","./fy.js":"7118","./ga":"5120","./ga.js":"5120","./gd":"f6b4","./gd.js":"f6b4","./gl":"8840","./gl.js":"8840","./gom-deva":"aaf2","./gom-deva.js":"aaf2","./gom-latn":"0caa","./gom-latn.js":"0caa","./gu":"e0c5","./gu.js":"e0c5","./he":"c7aa","./he.js":"c7aa","./hi":"dc4d","./hi.js":"dc4d","./hr":"4ba9","./hr.js":"4ba9","./hu":"5b14","./hu.js":"5b14","./hy-am":"d6b6","./hy-am.js":"d6b6","./id":"5038","./id.js":"5038","./is":"0558","./is.js":"0558","./it":"6e98","./it-ch":"6f12","./it-ch.js":"6f12","./it.js":"6e98","./ja":"079e","./ja.js":"079e","./jv":"b540","./jv.js":"b540","./ka":"201b","./ka.js":"201b","./kk":"6d79","./kk.js":"6d79","./km":"e81d","./km.js":"e81d","./kn":"3e92","./kn.js":"3e92","./ko":"22f8","./ko.js":"22f8","./ku":"2421","./ku.js":"2421","./ky":"9609","./ky.js":"9609","./lb":"440c","./lb.js":"440c","./lo":"b29d","./lo.js":"b29d","./lt":"26f9","./lt.js":"26f9","./lv":"b97c","./lv.js":"b97c","./me":"293c","./me.js":"293c","./mi":"688b","./mi.js":"688b","./mk":"6909","./mk.js":"6909","./ml":"02fb","./ml.js":"02fb","./mn":"958b","./mn.js":"958b","./mr":"39bd","./mr.js":"39bd","./ms":"ebe4","./ms-my":"6403","./ms-my.js":"6403","./ms.js":"ebe4","./mt":"1b45","./mt.js":"1b45","./my":"8689","./my.js":"8689","./nb":"6ce3","./nb.js":"6ce3","./ne":"3a39","./ne.js":"3a39","./nl":"facd","./nl-be":"db29","./nl-be.js":"db29","./nl.js":"facd","./nn":"b84c","./nn.js":"b84c","./oc-lnc":"167b","./oc-lnc.js":"167b","./pa-in":"f3ff","./pa-in.js":"f3ff","./pl":"8d57","./pl.js":"8d57","./pt":"f260","./pt-br":"d2d4","./pt-br.js":"d2d4","./pt.js":"f260","./ro":"972c","./ro.js":"972c","./ru":"957c","./ru.js":"957c","./sd":"6784","./sd.js":"6784","./se":"ffff","./se.js":"ffff","./si":"eda5","./si.js":"eda5","./sk":"7be6","./sk.js":"7be6","./sl":"8155","./sl.js":"8155","./sq":"c8f3","./sq.js":"c8f3","./sr":"cf1e","./sr-cyrl":"13e9","./sr-cyrl.js":"13e9","./sr.js":"cf1e","./ss":"52bd","./ss.js":"52bd","./sv":"5fbd","./sv.js":"5fbd","./sw":"74dc","./sw.js":"74dc","./ta":"3de5","./ta.js":"3de5","./te":"5cbb","./te.js":"5cbb","./tet":"576c","./tet.js":"576c","./tg":"3b1b","./tg.js":"3b1b","./th":"10e8","./th.js":"10e8","./tk":"5aff","./tk.js":"5aff","./tl-ph":"0f38","./tl-ph.js":"0f38","./tlh":"cf755","./tlh.js":"cf755","./tr":"0e81","./tr.js":"0e81","./tzl":"cf51","./tzl.js":"cf51","./tzm":"c109","./tzm-latn":"b53d","./tzm-latn.js":"b53d","./tzm.js":"c109","./ug-cn":"6117","./ug-cn.js":"6117","./uk":"ada2","./uk.js":"ada2","./ur":"5294","./ur.js":"5294","./uz":"2e8c","./uz-latn":"010e","./uz-latn.js":"010e","./uz.js":"2e8c","./vi":"2921","./vi.js":"2921","./x-pseudo":"fd7e","./x-pseudo.js":"fd7e","./yo":"7f33","./yo.js":"7f33","./zh-cn":"5c3a","./zh-cn.js":"5c3a","./zh-hk":"49ab","./zh-hk.js":"49ab","./zh-mo":"3a6c","./zh-mo.js":"3a6c","./zh-tw":"90ea","./zh-tw.js":"90ea"};function n(s){var e=r(s);return a(e)}function r(s){if(!a.o(t,s)){var e=new Error("Cannot find module '"+s+"'");throw e.code="MODULE_NOT_FOUND",e}return t[s]}n.keys=function(){return Object.keys(t)},n.resolve=r,s.exports=n,n.id="4678"},"4a00":function(s,e,a){},"56d7":function(s,e,a){"use strict";a.r(e);a("e260"),a("e6cf"),a("cca6"),a("a79d");var t,n,r=a("2b0e"),c=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{attrs:{id:"app"}},[a("Header"),a("div",{staticClass:"container-fluid"},[a("SensorGroup",{attrs:{groups:s.groups,timestamp:s.timestamp}})],1)],1)},i=[],l=a("bc3a"),o=a.n(l),d=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("header",{staticClass:"mb-3"},[a("nav",{staticClass:"navbar navbar-expand-lg navbar-dark bg-dark"},[a("div",{staticClass:"container-fluid"},[a("a",{staticClass:"navbar-brand nav",attrs:{href:"/"}},[s._v(s._s(s.appName))]),s._m(0),a("div",{staticClass:"collapse navbar-collapse",attrs:{id:"navbarNav"}},[a("ul",{staticClass:"navbar-nav"},s._l(s.menuItems,(function(e){return a("li",{staticClass:"nav-item"},[a("a",{staticClass:"nav-link",class:{active:e.isActive},attrs:{"aria-current":"page",href:e.href}},[s._v(s._s(e.name))])])})),0)])])])])},u=[function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("button",{staticClass:"navbar-toggler",attrs:{type:"button","data-bs-toggle":"collapse","data-bs-target":"#navbarNav","aria-controls":"navbarNav","aria-expanded":"false","aria-label":"Toggle navigation"}},[a("span",{staticClass:"navbar-toggler-icon"})])}],f={name:"Header",data:function(){return{appName:"Sensors Monitor",menuItems:[{name:"Temperature",isActive:!0,href:"#"}]}}},b=f,j=a("2877"),p=Object(j["a"])(b,d,u,!1,null,null,null),m=p.exports,v=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{staticClass:"d-grid"},s._l(s.groups,(function(e){return e.sensors.length>0?a("div",{staticClass:"card border-dark mb-3"},[a("div",{staticClass:"card-header border-dark"},[s._v(s._s(e.name))]),a("div",{staticClass:"card-body"},[a("div",{staticClass:"row row-cols-5"},s._l(e.sensors,(function(e){return a("div",{staticClass:"col"},[a("Sensor",{attrs:{name:e.name,timestamp:s.timestamp,value:e.value}})],1)})),0)])]):s._e()})),0)},h=[],g=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{staticClass:"py-2"},[a("div",{staticClass:"card"},[a("div",{staticClass:"card-header bg-white border-0"},[a("span",{staticClass:"card-title"},[s._v(s._s(s.name))]),a("div",{staticClass:"card-tools"},[a("span",[a("span",{staticClass:"badge text-white",class:s.indicatorClass()},[s._v(s._s(s.value)+"°С")])])])]),a("div",{staticClass:"card-body"},[a("SensorChart",{attrs:{"chart-id":s.chartId,height:196,timestamp:s.timestamp,value:s.value}})],1)])])},y=[],k=(a("fb6a"),a("1fca")),_={extends:k["a"],name:"SensorChart",props:["value","timestamp"],mixins:[k["b"].reactiveData],data:function(){return{labelsStorage:[],valuesStorage:[],options:{elements:{point:{radius:0}},legend:{display:!1},plugins:{datalabels:{display:!1}},scales:{xAxes:[{display:!1}],yAxes:[{ticks:{suggestedMin:20,suggestedMax:100}}]},tooltips:{enabled:!1}}}},mounted:function(){this.renderChart(this.chartData,this.options)},watch:{timestamp:function(){var s=this.valuesStorage;s.push(this.value);var e=this.labelsStorage;e.push(this.valuesStorage.length);var a=this.valuesStorage.length-60;a>0&&(s=this.valuesStorage.slice(a),e=this.labelsStorage.slice(a)),this.valuesStorage=s,this.labelsStorage=e,this.chartData={labels:this.labelsStorage,datasets:[{borderColor:"#17a2b8",data:this.valuesStorage,backgroundColor:"#b1ecfa"}]}}}},C=_,S=Object(j["a"])(C,t,n,!1,null,null,null),w=S.exports,x={name:"Sensor",components:{SensorChart:w},props:["name","value","timestamp"],data:function(){return{chartId:null}},methods:{indicatorClass:function(){return this.value<60?"bg-success":this.value<75?"bg-warning":"bg-danger"}},mounted:function(){this.chartId="chart"+this._uid}},z=x,O=(a("57fc"),Object(j["a"])(z,g,y,!1,null,"76c438c3",null)),M=O.exports,A={name:"SensorGroup",props:["groups","timestamp"],components:{Sensor:M}},E=A,N=Object(j["a"])(E,v,h,!1,null,null,null),P=N.exports,I={name:"App",components:{Header:m,SensorGroup:P},data:function(){return{groups:[],timestamp:0}},methods:{updateSensors:function(){var s=this;o.a.get("/api/v1/indicators/status").then((function(e){e.data.success&&Array.isArray(e.data.result.indications)&&(s.groups=e.data.result.indications,s.timestamp=e.data.result.timestamp)})).catch((function(s){console.log(s)}))}},mounted:function(){var s=this;setInterval((function(){s.updateSensors()}),3e3)}},T=I,$=(a("034f"),Object(j["a"])(T,c,i,!1,null,null,null)),D=$.exports,G=a("5f5b"),H=a("b1e0");a("f9e3"),a("2dd8");r["default"].config.productionTip=!1,r["default"].use(G["a"]),r["default"].use(H["a"]),new r["default"]({render:function(s){return s(D)}}).$mount("#app")},"57fc":function(s,e,a){"use strict";a("4a00")},"85ec":function(s,e,a){}});