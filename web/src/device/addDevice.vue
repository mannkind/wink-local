<template>
    <div>
        <div class="input-group">
            <select v-model="radio" :class="{ disabled: progress != 0 }"
                class="form-control form-control-sm">
                <option value="" selected="selected">Select a Radio</option>
                <option v-for="radio in availableRadios" :value="radio">
                    {{ radio }}
                </option>
            </select>
            <span class="input-group-btn">
                <a @click.stop="addDevice()" 
                    :class="{ disabled: progress != 0 }"
                    tabindex="-1" class="btn btn-sm btn-primary">Add</a>
            </span>
        </div>
        <div>
            <progress v-if="progress != 0" :value="progress" 
                class="progress" max="60" />
        </div>
    </div>
</template>

<script>
module.exports = require('./addDevice.ts').default.vueComponentOptions
</script>

<style scoped>
progress { margin-top: 1rem; }
</style>
