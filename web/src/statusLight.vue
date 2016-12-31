<template>
	<div>
		<div class="row">
			<div class="col-md-12">
				<h2>Status Light Manager</h2>
				<div class="card">
					<div class="card-block">
						<sketch v-model="colors" @change-color="change" />
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
module.exports = require('./statusLight.ts').default.vueComponentOptions
</script>

<style scoped>
	div.card-block div.vue-color__sketch {
		width: auto;
		margin: auto;
	}
</style>