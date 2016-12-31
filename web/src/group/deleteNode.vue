<template>
	<div class="row">
		<div class="col col-md-4" v-for="node in group.Nodes">
			<div class="input-group input-group-sm">
				<span class="input-group-btn">
					<a tabindex="-1" class="btn btn-sm btn-danger" @click.stop="deleteNode(node)"
						:class="{ disabled: group.ActionInProgress }">Delete</a>
				</span>
				<input :value="node.Name" 
					:class="{ disabled: group.ActionInProgress }"
					class="form-control form-control-sm" 
					readonly="readonly" type="text" />
			</div>
		</div>
	</div>
</template>

<script>
module.exports = require('./deleteNode.ts').default.vueComponentOptions
</script>

<style scoped>
.col { margin-bottom: 1rem; }
</style>
