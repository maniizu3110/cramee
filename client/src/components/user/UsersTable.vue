<template>
  <div class="w-full">
    <div class="d-flex flex-column flex-grow-1">
      <v-card>
        <v-data-table
          :headers="headers"
          :items="users"
          class="flex-grow-1"
          @click:row="choose"
        >
        </v-data-table>
      </v-card>
    </div>
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
      ],
    };
  },
  methods: {
    choose(e) {
      this.$router.push(`/student/schedule/teachers/${e.ID}`);
    },
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
