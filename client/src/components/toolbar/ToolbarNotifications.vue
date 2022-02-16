<template>
  <v-menu offset-y left transition="slide-y-transition">
    <template v-slot:activator="{ on }">
      <v-badge bordered :content="items.length" offset-x="22" offset-y="22">
        <v-btn icon v-on="on">
          <v-icon>mdi-bell-outline</v-icon>
        </v-btn>
      </v-badge>
    </template>
    <v-card>
      <v-list three-line dense max-width="400">
        <v-subheader class="pa-2 font-weight-bold">Notifications</v-subheader>
        <div v-for="(item, index) in items" :key="index">
          <v-divider v-if="index > 0 && index < items.length" inset></v-divider>
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title
                v-text="item.status == 'pending' ? 'リクエスト' : ''"
              ></v-list-item-title>
              <v-list-item-subtitle
                class="caption"
                v-text="
                  `${moment(item.start_time).format('YYYY年M月D日H時M分')}` +
                  '〜' +
                  `${moment(item.end_time).format('YYYY年M月D日H時M分')}`
                "
              ></v-list-item-subtitle>
              <div class="float-right">
                <v-btn @click="updateSchedule(item.ID, true)" small>承認</v-btn>
                <v-btn @click="updateSchedule(item.ID, false)" small
                  >拒否</v-btn
                >
              </div>
            </v-list-item-content>
          </v-list-item>
        </div>
      </v-list>
    </v-card>
  </v-menu>
</template>
<script>
import moment from "moment";
export default {
  props: {
    items: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      moment: moment,
      dialog: false,
    };
  },
  methods: {
    reload() {
      this.$router.go({ path: this.$router.currentRoute.path, force: true });
    },
    updateSchedule(id, accept) {
      this.$axios
        .put(`v1/lecture/${id}`, {
          Status: accept ? "reserved" : "empty",
        })
        .then(() => {
          this.reload();
        });
    },
  },
};
</script>
