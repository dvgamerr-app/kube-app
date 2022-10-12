const os = require('os')
const readline = require('readline')
const { spawn } = require('child_process')

let child
let fails = 0
let goBinary = './kube-app' //or template.exe

function setPage(html) {
	const app = document.getElementById('app')
	app.innerHTML = html
	//set focus for autofocus element
	const elem = document.querySelector('input[autofocus]')
	if (elem != null) elem.focus()
}

const bodyMessage = msg => {
	setPage(`'<h1>${msg}</h1>'`)
}

const startBinarayGO = async () => {
	bodyMessage('Loading...')
	await new Promise()
	child = spawn(goBinary, { maxBuffer: 1024 * 500 })

	const rl = readline.createInterface({
		input: child.stdout
	})

	rl.on('line', (data) => {
		console.log(`Received: ${data}`)

		if (data.charAt(0) == '$') {
			eval(data.substring(1))
		} else {
			setPage(data)
		}
	})

	child.stderr.on('data', (data) => {
		console.log(`stderr: ${data}`)
	})

	child.on('close', (code) => {
		bodyMessage(`process exited with code ${code}`)
		restart_process()
	})

	child.on('error', err => {
		bodyMessage(`Failed '${err}' to start child process.`)
		restart_process()
	})
}

function restart_process() {
	setTimeout(function () {
		fails++
		if (fails > 5) {
			close()
		} else {
			startBinarayGO()
		}
	}, 5000)
}

// function element_as_object(elem) {
// 	var obj = {
// 			properties: {}
// 	}
// 	for (var j = 0; j < elem.attributes.length; j++) {
// 			obj.properties[elem.attributes[j].name] = elem.attributes[j].value
// 	}
// 	//overwrite attributes with properties
// 	if (elem.value != null) {
// 			obj.properties['value'] = elem.value.toString()
// 	}
// 	if (elem.checked != null && elem.checked) {
// 			obj.properties['checked'] = 'true'
// 	} else {
// 			delete (obj.properties['checked'])
// 	}
// 	return obj
// }

// function element_by_tag_as_array(tag) {
// 	var items = []
// 	var elems = document.getElementsByTagName(tag)
// 	for (var i = 0; i < elems.length; i++) {
// 		items.push(element_as_object(elems[i]))
// 	}
// 	return items
// }

// function fire_event(name, sender) {
// 	var msg = {
// 		name: name,
// 		sender: element_as_object(sender),
// 		inputs: element_by_tag_as_array('input').concat(element_by_tag_as_array('select'))
// 	}
// 	child.stdin.write(JSON.stringify(msg))
// 	console.log(JSON.stringify(msg))
// }

// function fire_keypressed_event(e, keycode, name, sender) {
// 	if (e.keyCode === keycode) {
// 		e.preventDefault()
// 		fire_event(name, sender)
// 	}
// }

const avoid_reload = () => {
	if (sessionStorage.getItem('loaded') == 'true') {
		alert('go-webkit will fail when page reload. avoid using <form> or submit.')
		close()
	}
	sessionStorage.setItem('loaded', 'true')
}

if (os.platform().isWindows) {
	goBinary += '.exe'
}

avoid_reload()
startBinarayGO().catch(ex => {
	alert(ex)
	close()
})
