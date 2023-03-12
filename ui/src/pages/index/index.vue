<template>
  <view class="content">
    <view class="input-content">
      <view class="row">
        <text>占者：</text>
        <view class="item">
          <label for="yearOfBirth">生年：</label>
          <input
            class="input"
            id="yearOfBirth"
            type="number"
            v-model="yearOfBirth"
          />
        </view>

        <view class="item">
          <text>性别：</text>
          <radio-group @change="sexChange">
            <radio value="male" :checked="sex == 'male'" />
            <text>男</text>
            <radio value="female" :checked="sex == 'female'" />
            <text>女</text>
          </radio-group>
        </view>
      </view>

      <view class="row">
        <label>时间：</label>
        <view class="item">
          <input class="input" v-model="year" id="year" type="number" />
          <label for="year">年</label>
        </view>

        <view class="item">
          <input class="small-input" id="month" type="number" v-model="month" />
          <label for="month">月</label>
        </view>

        <view class="item">
          <input class="small-input" id="day" type="number" v-model="day" />
          <label for="day">日</label>
        </view>

        <view class="item">
          <input class="small-input" id="hour" type="number" v-model="hour" />
          <label for="hour">时</label>
        </view>

        <view class="item">
          <input
            class="small-input"
            id="minute"
            type="number"
            v-model="minute"
          />
          <label for="minute">分</label>
        </view>

        <view class="item">
          <input
            class="small-input"
            id="second"
            type="number"
            v-model="second"
          />
          <label for="second">秒</label>
        </view>

        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          fill="currentColor"
          class="bi bi-clock"
          viewBox="0 0 16 16"
          @click="toNowDateTime"
        >
          <path
            d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71V3.5z"
          />
          <path
            d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zm7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0z"
          />
        </svg>
      </view>
      <view class="row">
        <view class="item">
          <label>月将：</label>
          <picker
            @change="yueJiangChange"
            :value="yueJiangIndex"
            :range="yueJiangList"
          >
            <view>{{ yueJiangList[yueJiangIndex] }}</view>
          </picker>
        </view>

        <view class="item">
          <label>占时：</label>
          <picker
            @change="divinationTimeChange"
            :value="divinationTimeIndex"
            :range="diZhi"
          >
            <view>{{ diZhi[divinationTimeIndex] }}</view>
          </picker>
        </view>

        <view class="item">
          <label>昼夜：</label>
          <radio-group @change="diurnalChange">
            <radio value="1" :checked="diurnal == 1" />
            <text>昼占</text>
            <radio value="0" :checked="diurnal == 0" />
            <text>夜占</text>
          </radio-group>
        </view>
      </view>

      <navigator
        :url="
          '/pages/daliuren/pan?item=' +
          encodeURIComponent(JSON.stringify(requestData))
        "
        open-type="navigate"
        hover-class="navigator-hover"
      >
        <button type="primary" :loading="false" hover-class="button-hover">
          提交
        </button>
      </navigator>
    </view>
  </view>
</template>

<script setup lang="ts">
import { diZhi } from "@/constdef";
import type { RequestData } from "@/type";
import { nowDateTime } from "@/utils";
import type { request } from "@dcloudio/types/uni-app/uni/base/request";
import { computed, ref } from "vue";

const t = nowDateTime();
const yearOfBirth = ref(t.year);
const sex = ref("male");

const year = ref(t.year);
const month = ref(t.month);
const day = ref(t.day);
const hour = ref(t.hour);
const minute = ref(t.minute);
const second = ref(t.second);

const yueJiangList = ["默认"].concat(...diZhi);
const yueJiangIndex = ref(0);
const n = Math.floor((hour.value + 1) / 2) % 12;
const divinationTimeIndex = ref(n);
const diurnal = ref(1);
if (hour.value > 17 || hour.value < 7) diurnal.value = 0;

const requestData = computed<RequestData>(() => {
  return {
    year: Number(year.value),
    month: Number(month.value),
    day: Number(day.value),
    hour: Number(hour.value),
    minute: Number(minute.value),
    second: Number(second.value),
    yue_jiang:
      yueJiangIndex.value == 0 ? null : yueJiangList[yueJiangIndex.value],
    divination_time: diZhi[divinationTimeIndex.value],
    diurnal: diurnal.value == 1,
    year_of_birth: Number(yearOfBirth.value),
    masculine: sex.value == "male",
  };
});

function sexChange(event: { detail: { value: string } }) {
  sex.value = event.detail.value;
}

function yueJiangChange(e: { detail: { value: number } }) {
  yueJiangIndex.value = e.detail.value;
}

function divinationTimeChange(e: { detail: { value: number } }) {
  divinationTimeIndex.value = e.detail.value;
}

function diurnalChange(e: { detail: { value: number } }) {
  diurnal.value = e.detail.value;
}

function toNowDateTime() {
  const t = nowDateTime();
  year.value = t.year;
  month.value = t.month;
  day.value = t.day;
  hour.value = t.hour;
  minute.value = t.minute;
  second.value = t.second;
  let n = Math.floor((hour.value + 1) / 2) % 12;
  divinationTimeIndex.value = n;
  diurnal.value = 1;
  if (hour.value > 17 || hour.value < 7) diurnal.value = 0;
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  flex-direction: row;
  // margin-left: 30rpx;
  // align-items: center;
  justify-content: center;
}

.input-content {
  display: flex;
  flex-direction: column;
}

.row {
  display: flex;
  margin-top: 20rpx;
  margin-bottom: 20rpx;
  .item {
    display: flex;
    margin-right: 25rpx;
  }
}

.input {
  // height: 50rpx;
  // padding: 15rpx 25rpx;
  // line-height:50rpx;
  // font-size:28rpx;
  background: #fff;
  width: 80rpx;
}

.small-input {
  width: 40rpx;
}
</style>
