<template>
  <div class="min-w-0 w-full">
    <v-card class="mt-2 pa-2">
      <v-row class="fill-height">
        <v-col>
          <v-sheet height="64">
            <v-toolbar flat>
              <v-btn outlined class="mr-4" @click="setToday"> Today </v-btn>
              <v-btn fab text small color="grey darken-2" @click="prev">
                <v-icon small>mdi-chevron-left</v-icon>
              </v-btn>
              <v-toolbar-title v-if="$refs.calendar">
                {{ $refs.calendar.title }}
              </v-toolbar-title>
              <v-btn fab text small color="grey darken-2" @click="next">
                <v-icon small>mdi-chevron-right</v-icon>
              </v-btn>
              <v-spacer></v-spacer>
            </v-toolbar>
          </v-sheet>
          <v-sheet height="600">
            <v-calendar
              ref="calendar"
              v-model="focus"
              color="primary"
              :events="events"
              :event-color="getEventColor"
              :type="type"
              @click:more="viewDay"
              @click:date="viewDay"
              @change="updateRange"
            ></v-calendar>
            <v-menu
              v-model="selectedOpen"
              :close-on-content-click="false"
              :activator="selectedElement"
              offset-x
            >
              <v-card min-width="350px" flat>
                <v-toolbar :color="selectedEvent.color" dark>
                  <v-btn icon>
                    <v-icon>mdi-pencil</v-icon>
                  </v-btn>
                  <v-toolbar-title
                    v-html="selectedEvent.name"
                  ></v-toolbar-title>
                  <v-spacer></v-spacer>
                  <v-btn icon>
                    <v-icon>mdi-heart</v-icon>
                  </v-btn>
                  <v-btn icon>
                    <v-icon>mdi-dots-vertical</v-icon>
                  </v-btn>
                </v-toolbar>
                <v-card-text>
                  <span v-html="selectedEvent.details"></span>
                </v-card-text>
                <v-card-actions>
                  <v-btn text color="secondary" @click="selectedOpen = false">
                    Cancel
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-menu>
          </v-sheet>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script>
//TODO:複雑になりすぎているのでリファクタリング
import moment from "moment";
import { mapState } from "vuex";
export default {
  data: () => ({
    picker: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
      .toISOString()
      .substr(0, 10),
    moment: moment,
    dialog: false,
    focus: "",
    type: "month",
    time: null,
    date: "",
    timePicker: {
      startTime: null,
      startTimeDialog: false,
      endTime: null,
      endTimeDialog: false,
    },
    typeToLabel: {
      month: "Month",
      week: "Week",
      day: "Day",
    },
    selectedEvent: {},
    selectedElement: null,
    selectedOpen: false,
    events: [],
    kind: {
      empty: { status: "空き", color: "grey darken-1" },
      pending: { status: "リクエスト済", color: "green" },
      reserved: { status: "予約済", color: "blue" },
      finished: { status: "完了", color: "green" },
      absent: { status: "欠席", color: "orange" },
    },
  }),
  computed: {
    ...mapState("auth", ["isStudent", "isLogin", "id"]),
  },
  mounted() {
    this.$refs.calendar.checkChange();
    this.date = moment().toISOString();
  },
  methods: {
    viewDay({ date }) {
      this.focus = date;
      this.type = "day";
      this.date = moment(date).toISOString();
    },
    getEventColor(event) {
      return event.color;
    },
    setToday() {
      this.focus = "";
    },
    prev() {
      this.$refs.calendar.prev();
    },
    next() {
      this.$refs.calendar.next();
    },

    updateRange({ start, end }) {
      //TODO:APIを叩いて登録してあるスケジュールをカレンダーに表示
      const min = new Date(`${start.date}T00:00:00`);
      const max = new Date(`${end.date}T23:59:59`);
      this.$axios
        .get("v1/student-lecture-schedule", {
          params: {
            Query: [
              `StartTime >= ${moment(min).toISOString()}`,
              `EndTime =< ${moment(max).toISOString()}`,
              `StudentID == ${this.id}`,
            ],
          },
        })
        .then((res) => {
          res.data.Data.forEach((el) => {
            //TODO:kindを状態によって変更する（emptyで固定されている）
            this.events.push({
              name: this.kind[el.status].status,
              start: moment(el.start_time).format("YYYY-MM-DD hh:mm"),
              end: moment(el.end_time).format("YYYY-MM-DD hh:mm"),
              color: this.kind[el.status].color,
            });
          });
        });
    },
  },
};
</script>
