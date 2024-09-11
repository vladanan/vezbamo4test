// const axios = require('axios')
import axios from 'axios';

let ts = new Date()
let ti = new Date()
let isum = 0
let calls0 = 0
let calls1 = 0
let ok = 0
let nok = 0
console.log("let's node:", ts)
let url = "http://192.168.0.226:10000/api/v/test/id/1?"
url = "http://192.168.0.226:10000/user_portal?"
url = "http://127.0.0.1:10000/custom_apis?"
url = "http://192.168.2.100:10000/custom_apis?"
url = "http://127.0.0.1:7331/custom_apis?"

for (let i = 0; i < 73; i++) {
	axios.get(`${url}call=${i}`)
	.then(function (res) {
		// let dif = new Date() - ti
		ok++
		isum = isum + dif
		// console.log(i, res.status)
		console.log(i, res.status, dif/1000, isum/1000, new Date(), calls1++)
		// ti = new Date()
	})
	.catch(function (error) {
		nok++
		if (typeof error.response != "undefined") {
			console.log(error.response.data)
		}
		// console.log(error)
		// if (i == 72) {
		// 	console.log("async done", new Date() - ts)
		// }
		// console.log("err", i	, calls0++)
		console.log("err", i, error.status, calls1, calls0++)
	});
}

console.log("node done:", new Date(), new Date() - ts, ok, nok)
