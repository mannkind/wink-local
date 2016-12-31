<template>
    <div class="row">
        <div class="col col-sm-4" v-for="(device, index) in devices">
            <div class="input-group input-group-sm">
                <span class="input-group-btn">
                    <a @click.stop="deleteDevice(device)"
                        :class="{ disabled: device.ActionInProgress }"
                        tabindex="-1" class="btn btn-sm btn-danger">Delete</a>
                </span>
                <span class="input-group-addon">
                    {{ device.ID }}</span>
                <input v-model="devices[index].Name" 
                    :class="{ disabled: device.ActionInProgress }"
                    class="form-control form-control-sm" 
                    name="name" type="text" />
                <span class="input-group-btn">
                    <a :class="{ disabled: device.ActionInProgress }"
                        @click.stop="updateDevice(device)" 
                        tabindex="-1" class="btn btn-sm btn-primary">Save</a>
                </span>
            </div> 
        </div>
    </div>
</template>

<script>
module.exports = require('./listDevice.ts').default.vueComponentOptions
</script>

<style scoped>
.col { margin-bottom: 1rem; }
</style>
