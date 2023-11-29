"use strict";(self.webpackChunktoolbox=self.webpackChunktoolbox||[]).push([[70],{3070:function(e,n,t){t.r(n);var o=t(9439),r=t(2791),i=t(7391),a=t(9164),c=t(6314),u=t(4294),s=t(184);n.default=function(){var e=(0,r.useState)(""),n=(0,o.Z)(e,2),t=n[0],l=n[1],d=(0,r.useState)(""),v=(0,o.Z)(d,2),f=v[0],p=v[1];return(0,s.jsx)("div",{children:(0,s.jsx)(a.Z,{fixed:!0,children:(0,s.jsxs)(c.Z,{spacing:3,children:[(0,s.jsxs)(c.Z,{spacing:2,children:[(0,s.jsx)(i.Z,{label:"Encode String",variant:"outlined",value:t,onChange:function(e){l(e.target.value)}}),(0,s.jsx)(u.Z,{variant:"contained",onClick:function(){return p(window.functions.encode(t))},children:"Encode Base64"})]}),(0,s.jsxs)(c.Z,{spacing:2,children:[(0,s.jsx)(i.Z,{label:"Decode String",variant:"outlined",value:f,onChange:function(e){p(e.target.value)}}),(0,s.jsx)(u.Z,{variant:"contained",onClick:function(){return l(window.functions.decode(f))},children:"Decode Base64"})]})]})})})}},6314:function(e,n,t){t.d(n,{Z:function(){return N}});var o=t(4942),r=t(3366),i=t(7462),a=t(2791),c=t(3733),u=t(2466),s=t(4419),l=t(1217),d=t(3457),v=t(7078),f=t(8519),p=t(5080),h=t(1184),m=t(5682),Z=t(184),g=["component","direction","spacing","divider","children","className","useFlexGap"],k=(0,p.Z)(),x=(0,d.Z)("div",{name:"MuiStack",slot:"Root",overridesResolver:function(e,n){return n.root}});function b(e){return(0,v.Z)({props:e,name:"MuiStack",defaultTheme:k})}function j(e,n){var t=a.Children.toArray(e).filter(Boolean);return t.reduce((function(e,o,r){return e.push(o),r<t.length-1&&e.push(a.cloneElement(n,{key:"separator-".concat(r)})),e}),[])}var y=function(e){var n=e.ownerState,t=e.theme,r=(0,i.Z)({display:"flex",flexDirection:"column"},(0,h.k9)({theme:t},(0,h.P$)({values:n.direction,breakpoints:t.breakpoints.values}),(function(e){return{flexDirection:e}})));if(n.spacing){var a=(0,m.hB)(t),c=Object.keys(t.breakpoints.values).reduce((function(e,t){return("object"===typeof n.spacing&&null!=n.spacing[t]||"object"===typeof n.direction&&null!=n.direction[t])&&(e[t]=!0),e}),{}),s=(0,h.P$)({values:n.direction,base:c}),l=(0,h.P$)({values:n.spacing,base:c});"object"===typeof s&&Object.keys(s).forEach((function(e,n,t){if(!s[e]){var o=n>0?s[t[n-1]]:"column";s[e]=o}}));r=(0,u.Z)(r,(0,h.k9)({theme:t},l,(function(e,t){return n.useFlexGap?{gap:(0,m.NA)(a,e)}:{"& > :not(style):not(style)":{margin:0},"& > :not(style) ~ :not(style)":(0,o.Z)({},"margin".concat((r=t?s[t]:n.direction,{row:"Left","row-reverse":"Right",column:"Top","column-reverse":"Bottom"}[r])),(0,m.NA)(a,e))};var r})))}return r=(0,h.dt)(t.breakpoints,r)};var S=t(6934),w=t(1402),C=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},n=e.createStyledComponent,t=void 0===n?x:n,o=e.useThemeProps,u=void 0===o?b:o,d=e.componentName,v=void 0===d?"MuiStack":d,p=t(y),h=a.forwardRef((function(e,n){var t=u(e),o=(0,f.Z)(t),a=o.component,d=void 0===a?"div":a,h=o.direction,m=void 0===h?"column":h,k=o.spacing,x=void 0===k?0:k,b=o.divider,y=o.children,S=o.className,w=o.useFlexGap,C=void 0!==w&&w,N=(0,r.Z)(o,g),P={direction:m,spacing:x,useFlexGap:C},R=(0,s.Z)({root:["root"]},(function(e){return(0,l.Z)(v,e)}),{});return(0,Z.jsx)(p,(0,i.Z)({as:d,ownerState:P,ref:n,className:(0,c.Z)(R.root,S)},N,{children:b?j(y,b):y}))}));return h}({createStyledComponent:(0,S.ZP)("div",{name:"MuiStack",slot:"Root",overridesResolver:function(e,n){return n.root}}),useThemeProps:function(e){return(0,w.Z)({props:e,name:"MuiStack"})}}),N=C}}]);
//# sourceMappingURL=70.a9933f87.chunk.js.map