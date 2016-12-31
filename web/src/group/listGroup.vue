<template>
    <div>
        <div v-for="(group, index) in groups">
            <div class="row">
                <div class="col-sm-4">
                    <div class="input-group input-group-sm">
                        <span class="input-group-btn">
                            <a @click.stop="deleteGroup(group)"
                                :class="{ disabled: group.ActionInProgress }"
                                tabindex="-1" class="btn btn-sm btn-danger">Delete</a>
                        </span>
                        <span class="input-group-addon">{{ group.ID }}</span>
                        <input :value="group.Name" 
                            :class="{ disabled: group.ActionInProgress }"
                            class="form-control form-control-sm" 
                            readonly="readonly" type="text" />
                        <span class="input-group-addon">
                            {{ group.Nodes.length }}</span>
                    </div>
                </div>
                <div class="col-sm-8">
                    <add-node :group="group" :devices="devices" />
                </div>
            </div>
            <div class="row">
                <div class="col-sm-2">
                    <strong>Assoc. Devices</strong>
                </div>
                <div class="col-sm-10">
                    <delete-node :group="group" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
module.exports = require('./listGroup.ts').default.vueComponentOptions
</script>

<style scoped>
.row { margin-bottom: 1rem; }
</style>
