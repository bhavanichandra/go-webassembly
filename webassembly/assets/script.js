const go = new Go();
WebAssembly.instantiateStreaming(
	fetch('javascriptgo.wasm'),
	go.importObject
).then((result) => {
	go.run(result.instance);
});
const textArea = document.getElementById('exampleFormControlTextarea1');
const form = document.getElementById('input-form');
const output = document.getElementById('jsonOut');

console.log('form', form);

form.addEventListener('submit', (event) => {
	event.preventDefault();
	const jsonStr = textArea.value;
	const json = JSONFormatter(jsonStr);
	output.textContent = json;
});
