(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["register"],{"73cf":function(e,s,r){"use strict";r.r(s);var t=function(){var e=this,s=e.$createElement,r=e._self._c||s;return r("div",{staticClass:"register"},[r("h1",[e._v("Register")]),r("form",{on:{submit:function(s){return s.preventDefault(),e.register(s)}}},[r("p",[r("label",{attrs:{for:"username"}},[e._v("Username")]),r("input",{directives:[{name:"model",rawName:"v-model",value:e.username,expression:"username"}],attrs:{id:"username",name:"username"},domProps:{value:e.username},on:{input:function(s){s.target.composing||(e.username=s.target.value)}}}),r("br"),r("label",{attrs:{for:"password"}},[e._v("Password")]),r("input",{directives:[{name:"model",rawName:"v-model",value:e.password,expression:"password"}],attrs:{id:"password",name:"password"},domProps:{value:e.password},on:{input:function(s){s.target.composing||(e.password=s.target.value)}}})]),r("input",{attrs:{type:"submit",value:"login"}})])])},a=[],n=r("852e"),o=r.n(n),u=r("5665"),i=r("bc3a"),m=r.n(i),p={data:function(){return{username:null,password:null}},methods:{register:function(){this.username||window.alert("Empty username"),this.password||window.alert("Empty password");var e=this;m.a.post(u["a"].postRegister,{username:this.username,password:this.password}).then((function(){o.a.set("nienna_username",e.username,{expires:30}),e.$store.commit("login",e.username),e.$router.push("/")})).catch((function(e){console.error("Catch error:",e)}))}}},l=p,d=r("2877"),c=Object(d["a"])(l,t,a,!1,null,null,null);s["default"]=c.exports}}]);
//# sourceMappingURL=register.504a09d4.js.map