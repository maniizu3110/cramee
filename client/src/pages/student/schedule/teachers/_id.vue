<template>
  <div class="w-full">
    <teacher-calendar :id="Number($route.params.id)" :editable="false" />
  </div>
</template>
<script>
import TeacherCalendar from "../../../../components/calendar/TeacherCalendar.vue";
export default {
  layout: "student",
  components: { TeacherCalendar },
};
</script>
