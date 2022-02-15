<template>
  <div v-if="id" class="w-full">
    <teacher-calender :id="id" />
  </div>
</template>
<script>
import { mapState } from "vuex";
import TeacherCalender from "../../components/calendar/TeacherCalendar.vue";
export default {
  components: {
    TeacherCalender,
  },
  computed: {
    ...mapState("auth", ["isStudent", "isLogin", "id"]),
  },
};
</script>
