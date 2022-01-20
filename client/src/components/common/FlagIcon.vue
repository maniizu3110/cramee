<template>
  <span class="flag-icon" :class="[`flag-icon-${flag}`, { 'flag-round': round }]"></span>
</template>

<script>
import 'flag-icon-css/sass/flag-icon.scss'
export default {
  props: {
    flag: {
      type: String,
      default: 'us'
    },
    round: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style lang="scss" scoped>
.flag-icon {
  height: 22px;
  width: 22px;

  &.flag-round {
    background-size: cover;
    border-radius: 100%;
    height: 26px;
    width: 26px;
  }
}
</style>
