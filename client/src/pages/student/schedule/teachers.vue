<template>
  <div v-if="users">
    <users-table :users="users" />
  </div>
</template>
<script>
import UsersTable from "../../../components/user/UsersTable.vue";

export default {
  components: {
    UsersTable,
  },
  data() {
    return {
      users: [],
    };
  },
  mounted() {
    this.$axios.get("v1/teacher").then((res) => {
      this.users = res.data.Data
    });
  },
};
</script>
