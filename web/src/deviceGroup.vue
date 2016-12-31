<template>
	<div>
		<div class="row">
			<div class="col-md-12">
				<h2>Device Manager</h2>
				<div class="card">
					<div class="card-block">
						<div class="row">
							<div class="col-md-4">
								<h3>New Device</h3>
								<add-device />
							</div>
						</div>
						<hr />
						<div class="row">
							<div class="col-md-12">
							
								<h3>Existing Devices</h3>
								<list-device :devices="devices" />
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="row">
			<div class="col-md-12">
				<h2>Group Manager</h2>
				<div class="card">
					<div class="card-block">
						<div class="row">
							<div class="col-md-4">
								<h3>New Group</h3>
								<add-group  />
							</div>
						</div>
						<hr />
						<h3>Existing Groups</h3>
						<list-group :groups="groups" :devices="devices" />
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
module.exports = require('./deviceGroup.ts').default.vueComponentOptions
</script>

<style scoped>

</style>
