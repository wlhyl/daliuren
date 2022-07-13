<template>
  <!-- <div class="home"> -->
  <h1>大六壬</h1>
  <div>
    <label>占测人：</label>
    <label>生年：</label>
    <el-input-number
      v-model="yearOfBirth"
      class="mx-4"
      :min="100"
      :max="9999"
      size="small"
      controls-position="right"
    />

    <label>性别：</label>
    <el-radio-group v-model="sex">
      <el-radio label="male">男</el-radio>
      <el-radio label="female">女</el-radio>
    </el-radio-group>
  </div>
  <div>
    <label>时间：</label>
    <el-input-number
      v-model="year"
      class="mx-4"
      :min="100"
      :max="9999"
      size="small"
      controls-position="right"
    />

    <label for="year">年</label>
    <el-input-number
      v-model="month"
      class="mx-4"
      :min="1"
      :max="12"
      size="small"
      controls-position="right"
    />

    <label for="month">月</label>
    <label>生年：</label>
    <el-input-number
      v-model="day"
      class="mx-4"
      :min="1"
      :max="31"
      size="small"
      controls-position="right"
    />
    <label for="day">日</label>
    <el-input-number
      v-model="hour"
      class="mx-4"
      :min="0"
      :max="23"
      size="small"
      controls-position="right"
    />

    <label for="hour">时</label>
    <el-input-number
      v-model="minute"
      class="mx-4"
      :min="0"
      :max="59"
      size="small"
      controls-position="right"
    />
    <label for="minute">分</label>
    <el-input-number
      v-model="second"
      class="mx-4"
      :min="0"
      :max="59"
      size="small"
      controls-position="right"
    />
    <label for="second">秒</label>
    <el-tooltip content="当前时间" placement="top">
      <el-button :icon="Clock" size="small" circle @click="toNowDateTime" />
    </el-tooltip>
  </div>
  <div>
    <label>月将：</label>
    <el-select v-model="yueJiang" class="m-2" placeholder="Select" size="small">
      <el-option label="默认" value=""></el-option>
      <el-option
        v-for="item in diZhi"
        :key="item"
        :label="item"
        :value="item"
      />
    </el-select>

    <label>占时：</label>
    <el-select
      v-model="divinationTime"
      class="m-2"
      placeholder="Select"
      size="small"
    >
      <el-option
        v-for="item in diZhi"
        :key="item"
        :label="item"
        :value="item"
      />
    </el-select>

    <label>昼夜：</label>
    <el-select v-model="diurnal" class="m-2" placeholder="Select" size="small">
      <el-option label="昼占" value="1" />
      <el-option label="夜占" value="0" />
    </el-select>
  </div>

  <el-button type="primary" @click="pan">提交</el-button>
  <!-- </div> -->
  <div id="err">
    <span v-if="shiPan.err">{{ shiPan.errMessage }}</span>
  </div>
  <div v-if="shiPan.completed">
    <div>
      农历：{{ shiPan.data.lunarCalendar.year }}年
      {{ shiPan.data.lunarCalendar.month }}
      {{ shiPan.data.lunarCalendar.day }}
      {{ shiPan.data.lunarCalendar.timeGanZhi }}时
    </div>
    <div>
      节气：
      {{ shiPan.data.solarTermFirst.name }}:
      {{ shiPan.data.solarTermFirst.year }}-
      {{ shiPan.data.solarTermFirst.month }}-
      {{ shiPan.data.solarTermFirst.day }}
      {{ shiPan.data.solarTermFirst.hour }}:
      {{ shiPan.data.solarTermFirst.minute }}:
      {{ shiPan.data.solarTermFirst.second }}

      {{ shiPan.data.solarTermSecond.name }}:
      {{ shiPan.data.solarTermSecond.year }}-
      {{ shiPan.data.solarTermSecond.month }}-
      {{ shiPan.data.solarTermSecond.day }}
      {{ shiPan.data.solarTermSecond.hour }}:
      {{ shiPan.data.solarTermSecond.minute }}:
      {{ shiPan.data.solarTermSecond.second }}
    </div>
    <div>
      干支：
      {{ shiPan.data.ganZhi.year }}年 {{ shiPan.data.ganZhi.month }}月
      {{ shiPan.data.ganZhi.day }}日 {{ shiPan.data.ganZhi.time }}时
    </div>
    <div>
      空亡：{{ shiPan.data.kongWang[0] }}，{{ shiPan.data.kongWang[1] }}
    </div>
    <div>
      本命：{{ shiPan.data.yearOfBirthGanZhi }} 行年：{{ shiPan.data.xingNian }}
    </div>

    <table id="sanchua" border="0px" cellspacing="5px">
      <tr>
        <td>{{ shiPan.data.sanChuan.liuQing[0] }}</td>
        <td>{{ shiPan.data.sanChuan.dunGan[0] }}</td>
        <td>{{ shiPan.data.sanChuan.chu }}</td>
        <td>{{ toTianJiang(shiPan.data.sanChuan.chu) }}</td>
      </tr>

      <tr>
        <td>{{ shiPan.data.sanChuan.liuQing[1] }}</td>
        <td>{{ shiPan.data.sanChuan.dunGan[1] }}</td>
        <td>{{ shiPan.data.sanChuan.zhong }}</td>
        <td>{{ toTianJiang(shiPan.data.sanChuan.zhong) }}</td>
      </tr>

      <tr>
        <td>{{ shiPan.data.sanChuan.liuQing[2] }}</td>
        <td>{{ shiPan.data.sanChuan.dunGan[2] }}</td>
        <td>{{ shiPan.data.sanChuan.mo }}</td>
        <td>{{ toTianJiang(shiPan.data.sanChuan.mo) }}</td>
      </tr>
    </table>

    <table id="sike" border="0px" cellspacing="10px">
      <tr>
        <td>{{ toTianJiang(shiPan.data.sike.zhiYing) }}</td>
        <td>{{ toTianJiang(shiPan.data.sike.zhiYang) }}</td>
        <td>{{ toTianJiang(shiPan.data.sike.ganYing) }}</td>
        <td>{{ toTianJiang(shiPan.data.sike.ganYang) }}</td>
      </tr>
      <tr>
        <td>{{ shiPan.data.sike.zhiYing }}</td>
        <td>{{ shiPan.data.sike.zhiYang }}</td>
        <td>{{ shiPan.data.sike.ganYing }}</td>
        <td>{{ shiPan.data.sike.ganYang }}</td>
      </tr>

      <tr>
        <td>{{ shiPan.data.sike.zhiYang }}</td>
        <td>{{ shiPan.data.sike.zhi }}</td>
        <td>{{ shiPan.data.sike.ganYang }}</td>
        <td>{{ shiPan.data.sike.gan }}</td>
      </tr>
    </table>

    <!-- 天盘 -->
    <table id="tian_pan" border="0px" cellspacing="10px">
      <tr>
        <td></td>
        <td>{{ toTianJiang(shiPan.data.tianPan[5]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[6]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[7]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[8]) }}</td>
        <td></td>
      </tr>
      <tr>
        <td></td>
        <td>{{ shiPan.data.tianPan[5] }}</td>
        <td>{{ shiPan.data.tianPan[6] }}</td>
        <td>{{ shiPan.data.tianPan[7] }}</td>
        <td>{{ shiPan.data.tianPan[8] }}</td>
        <td></td>
      </tr>
      <tr>
        <td>{{ toTianJiang(shiPan.data.tianPan[4]) }}</td>
        <td>{{ shiPan.data.tianPan[4] }}</td>
        <td></td>
        <td></td>
        <td>{{ shiPan.data.tianPan[9] }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[9]) }}</td>
      </tr>
      <tr>
        <td>{{ toTianJiang(shiPan.data.tianPan[3]) }}</td>
        <td>{{ shiPan.data.tianPan[3] }}</td>
        <td></td>
        <td></td>
        <td>{{ shiPan.data.tianPan[10] }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[10]) }}</td>
      </tr>
      <tr>
        <td></td>
        <td>{{ shiPan.data.tianPan[2] }}</td>
        <td>{{ shiPan.data.tianPan[1] }}</td>
        <td>{{ shiPan.data.tianPan[0] }}</td>
        <td>{{ shiPan.data.tianPan[11] }}</td>
        <td></td>
      </tr>
      <tr>
        <td></td>
        <td>{{ toTianJiang(shiPan.data.tianPan[2]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[1]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[0]) }}</td>
        <td>{{ toTianJiang(shiPan.data.tianPan[11]) }}</td>
        <td></td>
      </tr>
    </table>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { Clock } from "@element-plus/icons-vue";
import axios from "@/http-common";
import { DaliurenPan } from "@/type";

let t = nowDateTime();
let yearOfBirth = ref(t.year);
let sex = ref("male");

let year = ref(t.year);
let month = ref(t.month);
let day = ref(t.day);
let hour = ref(t.hour);
let minute = ref(t.minute);
let second = ref(t.second);

let diZhi = [
  "子",
  "丑",
  "寅",
  "卯",
  "辰",
  "巳",
  "午",
  "未",
  "申",
  "酉",
  "戌",
  "亥",
];
let yueJiang = ref("");

let n = Math.floor((hour.value + 1) / 2) % 12;
let divinationTime = ref(diZhi[n]);

let diurnal = ref("1");
if (hour.value > 17 || hour.value < 7) diurnal.value = "0";

function toNowDateTime() {
  t = nowDateTime();
  year.value = t.year;
  month.value = t.month;
  day.value = t.day;
  hour.value = t.hour;
  minute.value = t.minute;
  second.value = t.second;

  let n = Math.floor((hour.value + 1) / 2) % 12;
  divinationTime.value = diZhi[n];

  diurnal.value = "1";
  if (hour.value > 17 || hour.value < 7) diurnal.value = "0";
}

function nowDateTime(): {
  year: number;
  month: number;
  day: number;
  hour: number;
  minute: number;
  second: number;
} {
  let t = new Date();
  let year = t.getFullYear();
  let month = t.getMonth() + 1;
  let day = t.getDate();
  let hour = t.getHours();
  let minute = t.getMinutes();
  let second = t.getSeconds();
  return { year, month, day, hour, minute, second };
}

// 获取某地支的天将
function toTianJiang(d: string): string {
  let n = diZhi.indexOf(d);
  if (n == -1) return "";
  return shiPan.data.tianJiangPan[n];
}

// 从server获取式盘
let daliurenPan: DaliurenPan = {
  lunarCalendar: {
    year: "",
    month: "",
    day: "",
    timeGanZhi: "",
  },
  solarTermFirst: {
    name: "",
    year: 0,
    month: 0,
    day: 0,
    hour: 0,
    minute: 0,
    second: 0,
  },
  solarTermSecond: {
    name: "",
    year: 0,
    month: 0,
    day: 0,
    hour: 0,
    minute: 0,
    second: 0,
  },
  ganZhi: {
    year: "",
    month: "",
    day: "",
    time: "",
  },
  kongWang: [],
  tianPan: [],
  tianJiangPan: [],
  sike: {
    gan: "",
    ganYang: "",
    ganYing: "",
    zhi: "",
    zhiYang: "",
    zhiYing: "",
  },
  sanChuan: {
    chu: "",
    zhong: "",
    mo: "",
    dunGan: [],
    liuQing: [],
  },
  yearOfBirthGanZhi: "",
  xingNian: "",
};
let shiPan = reactive({
  data: daliurenPan,
  completed: false,
  err: false,
  errMessage: "",
});

function pan() {
  let requestData = {
    year: year.value,
    month: month.value,
    day: day.value,
    hour: hour.value,
    minute: minute.value,
    second: second.value,
    yue_jiang: yueJiang.value,
    divination_time: divinationTime.value,
    diurnal: diurnal.value == "1",
    year_of_birth: yearOfBirth.value,
    masculine: sex.value == "male",
  };
  shiPan.err = false;
  shiPan.completed = false;
  axios
    .post("/daliuren", requestData)
    .then((response) => {
      shiPan.data = response.data;
      shiPan.err = false;
      shiPan.completed = true;
    })
    .catch((err) => {
      shiPan.err = true;
      shiPan.completed = false;
      shiPan.errMessage = err.response;
    });
}
</script>

<style lang="scss" scoped>
table {
  margin: auto;
}
</style>