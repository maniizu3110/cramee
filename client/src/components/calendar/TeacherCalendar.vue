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
              <v-row justify="center">
                <v-dialog
                  v-if="editable"
                  v-model="dialog"
                  persistent
                  max-width="600px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn color="primary" dark v-bind="attrs" v-on="on">
                      Add
                    </v-btn>
                  </template>
                  <v-card>
                    <v-card-title>
                      <span class="headline">予定追加</span>
                    </v-card-title>
                    <v-card-text>
                      <v-container>
                        <v-row cols="12" sm="10">
                          <v-col>
                            <v-row justify="center">
                              <v-date-picker
                                v-model="picker"
                                full-width
                              ></v-date-picker>
                            </v-row>
                          </v-col>
                        </v-row>
                        <v-row>
                          <v-col cols="6">
                            <v-dialog
                              ref="startdialog"
                              v-model="timePicker.startTimeDialog"
                              :return-value.sync="timePicker.startTime"
                              persistent
                              width="290px"
                            >
                              <template v-slot:activator="{ on, attrs }">
                                <v-text-field
                                  v-model="timePicker.startTime"
                                  label="開始時間"
                                  prepend-icon="mdi-clock-time-four-outline"
                                  readonly
                                  v-bind="attrs"
                                  v-on="on"
                                ></v-text-field>
                              </template>
                              <v-time-picker
                                v-if="timePicker.startTimeDialog"
                                v-model="timePicker.startTime"
                                full-width
                              >
                                <v-spacer></v-spacer>
                                <v-btn
                                  text
                                  color="primary"
                                  @click="timePicker.startTimeDialog = false"
                                >
                                  Cancel
                                </v-btn>
                                <v-btn
                                  text
                                  color="primary"
                                  @click="
                                    $refs.startdialog.save(timePicker.startTime)
                                  "
                                >
                                  OK
                                </v-btn>
                              </v-time-picker>
                            </v-dialog>
                          </v-col>
                          <v-col cols="6">
                            <v-dialog
                              ref="enddialog"
                              v-model="timePicker.endTimeDialog"
                              :return-value.sync="timePicker.endTime"
                              persistent
                              width="290px"
                            >
                              <template v-slot:activator="{ on, attrs }">
                                <v-text-field
                                  v-model="timePicker.endTime"
                                  label="終了時間"
                                  prepend-icon="mdi-clock-time-four-outline"
                                  readonly
                                  v-bind="attrs"
                                  v-on="on"
                                ></v-text-field>
                              </template>
                              <v-time-picker
                                v-if="timePicker.endTimeDialog"
                                v-model="timePicker.endTime"
                                full-width
                              >
                                <v-spacer></v-spacer>
                                <v-btn
                                  text
                                  color="primary"
                                  @click="timePicker.endTimeDialog = false"
                                >
                                  Cancel
                                </v-btn>
                                <v-btn
                                  text
                                  color="primary"
                                  @click="
                                    $refs.enddialog.save(timePicker.endTime)
                                  "
                                >
                                  OK
                                </v-btn>
                              </v-time-picker>
                            </v-dialog>
                          </v-col>
                        </v-row>
                      </v-container>
                    </v-card-text>
                    <v-card-actions>
                      <v-spacer></v-spacer>
                      <v-btn color="blue darken-1" text @click="dialog = false"
                        >Close</v-btn
                      >
                      <v-btn color="blue darken-1" text @click="saveSchedule"
                        >Save</v-btn
                      >
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-row>

              <v-menu bottom right>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn outlined v-bind="attrs" v-on="on">
                    <span>{{ typeToLabel[type] }}</span>
                    <v-icon right>mdi-menu-down</v-icon>
                  </v-btn>
                </template>
                <v-list>
                  <v-list-item @click="type = 'day'">
                    <v-list-item-title>Day</v-list-item-title>
                  </v-list-item>
                  <v-list-item @click="type = 'week'">
                    <v-list-item-title>Week</v-list-item-title>
                  </v-list-item>
                  <v-list-item @click="type = 'month'">
                    <v-list-item-title>Month</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-toolbar>
          </v-sheet>
          <v-sheet height="600">
            <v-calendar
              ref="calendar"
              v-model="focus"
              @click:event="showEvent"
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
                  <v-toolbar-title
                    v-html="selectedEvent.name"
                  ></v-toolbar-title>
                  <v-spacer></v-spacer>
                  <!--<v-btn icon>
                    <v-icon>mdi-heart</v-icon>
                  </v-btn>
                  <v-btn icon>
                    <v-icon>mdi-dots-vertical</v-icon>
                  </v-btn>-->
                </v-toolbar>
                <v-card-text>
                  <span
                    >{{ selectedEvent.start }}〜{{ selectedEvent.end }}</span
                  >
                </v-card-text>
                <v-card-actions>
                  <v-btn text color="secondary" @click="selectedOpen = false">
                    閉じる
                  </v-btn>
                  <v-btn text color="secondary" @click="deleteSchedule">
                    削除
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
export default {
  props: {
    id: {
      required: true,
    },
    editable: {
      type: Boolean,
      required: true,
      default: false,
    },
  },
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
  mounted() {
    this.$refs.calendar.checkChange();
    this.date = moment().toISOString();
  },
  methods: {
    reload() {
      this.$router.go({ path: this.$router.currentRoute.path, force: true });
    },
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
    showEvent({ nativeEvent, event }) {
      const open = () => {
        this.selectedEvent = event;
        this.selectedElement = nativeEvent.target;
        requestAnimationFrame(() =>
          requestAnimationFrame(() => (this.selectedOpen = true))
        );
      };

      if (this.selectedOpen) {
        this.selectedOpen = false;
        requestAnimationFrame(() => requestAnimationFrame(() => open()));
      } else {
        open();
      }

      nativeEvent.stopPropagation();
    },
    saveSchedule() {
      let pickerDate = moment(this.picker).format("YYYY-MM-DD");
      let start_time = moment(
        pickerDate + " " + this.timePicker.startTime
      ).toISOString();
      const end_time = moment(
        pickerDate + " " + this.timePicker.endTime
      ).toISOString();
      this.$axios
        .post("v1/lecture", {
          teacher_id: this.id,
          start_time: start_time,
          end_time: end_time,
        })
        .then((res) => {
          this.dialog = false;
        });
    },
    deleteSchedule(e) {
      this.$axios.delete(`v1/lecture/${this.selectedEvent.id}`).then((res) => {
        this.reload();
      });
    },
    updateRange({ start, end }) {
      //TODO:APIを叩いて登録してあるスケジュールをカレンダーに表示
      this.events = [];
      const min = new Date(`${start.date}T00:00:00`);
      const max = new Date(`${end.date}T23:59:59`);
      this.$axios
        .get("v1/lecture", {
          params: {
            Query: [
              `StartTime >= ${moment(min).toISOString()}`,
              `EndTime =< ${moment(max).toISOString()}`,
              `TeacherID == ${this.id}`,
            ],
            IncludeDeleted: false,
          },
        })
        .then((res) => {
          (res);
          res.data.Data.forEach((el) => {
            el.id = el.ID;
            el.name = this.kind[el.status].status;
            el.start = moment(el.start_time).format("YYYY-MM-DD hh:mm");
            el.end = moment(el.end_time).format("YYYY-MM-DD hh:mm");
            el.color = this.kind[el.status].color;
            this.events.push(el);
          });
        });
    },
  },
};
</script>
