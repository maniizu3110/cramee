<template>
  <div></div>
</template>
<script>
import { mapActions } from "vuex";
export default {
  async mounted() {
    await this.logout().then(() => this.$router.push("/"));
  },
  methods: {
    ...mapActions("auth", ["logout"]),
  },
};
</script>
