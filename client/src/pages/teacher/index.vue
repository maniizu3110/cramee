<template>
  <div class="w-full">
    <teacher-calender :id="id" :editable="true" />
  </div>
</template>
<script>
import { mapState } from "vuex";
import TeacherCalender from "../../components/calendar/TeacherCalendar.vue";
export default {
  layout: 'teacher',
  components: {
    TeacherCalender,
  },
  computed: {
    ...mapState("auth", ["isStudent", "isLogin", "id"]),
  },
};
</script>
