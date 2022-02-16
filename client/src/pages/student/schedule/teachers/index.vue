<template>
  <div class="w-full" v-if="users">
    <users-table :users="users" :editable="false" />
  </div>
</template>
<script>
import UsersTable from "../../../../components/user/UsersTable.vue";

export default {
  layout:'student',
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
      this.users = res.data.Data;
    });
  },
};
</script>
