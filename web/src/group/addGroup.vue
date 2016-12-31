<template>
	<div>
        <div class="input-group">
			<input class="form-control form-control-sm" type="text"
				placeholder="Enter a group name" selected="selected"
				v-model="name" :class="{ disabled: actionInProgress }" />
            <span class="input-group-btn">
				<a tabindex="-1" class="btn btn-sm btn-primary" 
					@click.stop="addGroup()" 
					:class="{ disabled: actionInProgress }">Add</a>
            </span>
        </div>
	</div>
</template>

<script>
module.exports = require('./addGroup.ts').default.vueComponentOptions
</script>

<style scoped>

</style>