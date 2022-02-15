<template>
  <div class="w-full">
    <student-calender />
  </div>
</template>
<script>
import StudentCalender from "../../components/calendar/StudentCalendar.vue";
export default {
  layout: "student",
  components: {
    StudentCalender,
  },
};
</script>
