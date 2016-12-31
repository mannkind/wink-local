<template>
	<div>
		<h1>{{ title }}</h1>
		<div class="row">
			<div class="col-md-12">
				<device-group />
			</div>
		</div>
		<div class="row">
			<div class="col-md-12">
				<status-light />
			</div>
		</div>
	</div>
</template>

<script>
module.exports = require('./app.ts').default.vueComponentOptions
</script>

<style scoped>
body {
	font-family: Verdana, Helvetica, sans-serif;
}
</style>
