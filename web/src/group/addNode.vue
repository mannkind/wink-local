<template>
	<div class="input-group">
		<select v-model="nodeId" class="form-control form-control-sm">
			<option value="" selected="selected">Select a Device</option>
			<option v-for="device in devices" :value="device.ID">
				{{ device.Name }}
			</option>
		</select>
		<span class="input-group-btn">
			<a tabindex="-1" class="btn btn-sm btn-primary" 
				@click.stop="addNode()" 
				:class="{ disabled: group.ActionInProgress }">Add</a>
		</span>
	</div>
</template>

<script>
module.exports = require('./addNode.ts').default.vueComponentOptions
</script>

<style scoped>

</style>
