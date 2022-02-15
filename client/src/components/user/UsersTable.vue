<template>
  <div class="d-flex flex-column flex-grow-1 full-width">
    <v-card>
      <v-data-table
        v-model="selectedUsers"
        :headers="headers"
        :items="users"
      >
      </v-data-table>
    </v-card>
  </div>
</template>

<script>
export default {
  props: {
    users: {
      type: Array,
      default: [],
    },
  },
  data() {
    return {
      isLoading: false,
      breadcrumbs: [
        {
          text: "Users",
          disabled: false,
          href: "#",
        },
        {
          text: "List",
        },
      ],

      searchQuery: "",
      selectedUsers: [],
      headers: [
        { text: "名前", align: "left", value: "name" },
        { text: "Email", value: "email" },
        { text: "登録日", value: "created_at" },
        //{ text: "最終ログイン", value: "lastSignIn" },
      ],
    };
  },
  watch: {
    selectedUsers(val) {},
  },
  methods: {
    searchUser() {},
    open() {},
  },
};
</script>

<style lang="scss" scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter,
.slide-fade-leave-to {
  transform: translateX(10px);
  opacity: 0;
}
</style>
