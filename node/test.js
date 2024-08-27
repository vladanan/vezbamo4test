const axios = require('axios')

console.log("node let's go")

for (let i = 990; i < 1000; i++) {
	axios.get(`http://192.168.2.100:10000/api/v/test/id/2?${i}`)
  .then(function (response) {
		console.log(response.data[0].Area)
	})
	.catch(function (error) {
		console.log(error)
	});
}
