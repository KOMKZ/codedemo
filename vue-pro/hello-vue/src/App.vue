<template lang="html">
	<div id="#app">
		<p>hello world</p>
	</div>
</template>

<script>
let routes, components = {
	index : {template : "<p>hello world</p>"},
	contact : {template : "<p>contact</p>"},
	detail : {template : "<p>detail</p>"},
	not_found : {template : "<p>not found</p>"}
}
export default {
	name : "app",

}
</script>

<style lang="css">
</style>
