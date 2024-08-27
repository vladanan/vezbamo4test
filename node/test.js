// const axios = require('axios')
import axios from 'axios';

let t = new Date()
let tt = new Date()
console.log("let's node:", t)
let url = "http://192.168.2.100:10000/api/v/test/id/2?"

for (let i = 990; i < 1000; i++) {
	axios.get(`${url}${i}`)
	.then(function (response) {
		console.log(i, new Date() - tt, response.data[0].Area)
		tt = new Date()
	})
	.catch(function (error) {
		console.log(error)
	});
}
console.log("node done:", new Date() - t)
