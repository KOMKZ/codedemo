/*

// 输出的时候是可以同时有 默认输出还有其他的东西，这不是冲突的
import * as app from './app.js';
console.log(app);

// 使用默认输入
import app  from "./app.js";
// app.js 中没有输出app这个变量，所以使用的默认app.js模块的默认输出, 注意使用默认输入不能使用{}， 否则会导致undefined
console.log(app);

// 将模块的变量输入到对象里面
import * as app from "./app.js";
console.log(app); // 此时可以认为app变量存在该模块所有输出的变量了

// 输入指定的变量
import {APP_INFO} from "./app.js";
console.log(APP_INFO);  // Object {name: "codedemo"}

// export的一些其他用法, 对模块的接口重新命名
export {app_info as my_app_info, app_config as my_app_config} from './app.js';
 */
