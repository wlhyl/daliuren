<template>
<view class="info-content">
  <view class="info">
    <!-- 占测人 -->
    <view class="item">
      <text>
        占者：{{ requestData.year_of_birth }}年{{
          requestData.masculine ? "男" : "女"
        }}
        </text>
    </view>
    <!-- 占测时间 -->
    <view class="item" >
        <text>
        时间：{{ requestData.year }}年{{ requestData.month }}月{{
          requestData.day
        }}日{{ requestData.hour }}:{{ requestData.minute }}:{{
          requestData.second
        }}
        </text>
    </view>
    
    <!-- 农历 -->
    <view class="item">
      <text
        >农历：{{ shiPan.data.lunar_calendar.year }}年{{
          shiPan.data.lunar_calendar.month
        }}{{ shiPan.data.lunar_calendar.day }}
        {{ shiPan.data.lunar_calendar.time_gan_zhi }}时
        </text>
    </view>
    <view class="item">
      <text
        >干支：{{ shiPan.data.si_zhu.year }}年{{ shiPan.data.si_zhu.month }}月{{
          shiPan.data.si_zhu.day
        }}日{{ shiPan.data.si_zhu.time }}时</text
      >
    </view>
<!-- 月将 -->
    <view>
      <text>
        月将：{{ shiPan.data.yue_jiang }}将 {{ requestData.divination_time }}时
        {{ requestData.diurnal ? "昼占" : "夜占" }}
        空亡：{{ shiPan.data.kong[0] }}、{{ shiPan.data.kong[1] }}</text
      >
    </view>

    <!-- 节气 -->
    <view class="item">
      <text>节气</text>：
      <view style="display: flex; flex-direction: column">
        <view>
          <text>
            {{ shiPan.data.solar_term_first.name }}:
            {{ shiPan.data.solar_term_first.year }}-{{
              shiPan.data.solar_term_first.month
            }}-{{ shiPan.data.solar_term_first.day }}
            {{ shiPan.data.solar_term_first.hour }}:{{
              shiPan.data.solar_term_first.minute
            }}:{{ shiPan.data.solar_term_first.second }}
          </text></view
        >
        <view
          ><text>
            {{ shiPan.data.solar_term_second.name }}:
            {{ shiPan.data.solar_term_second.year }}-{{
              shiPan.data.solar_term_second.month
            }}-{{ shiPan.data.solar_term_second.day }}
            {{ shiPan.data.solar_term_second.hour }}:{{
              shiPan.data.solar_term_second.minute
            }}:{{ shiPan.data.solar_term_second.second }}
          </text>
        </view>
      </view>
    </view>
    <!-- 行年 -->
    <view class="item">
      <text
        >本命：{{ shiPan.data.year_of_birth_gan_zhi }} 行年：{{
          shiPan.data.xing_nian
        }}</text
      >
    </view>
  </view>
</view>

  <!-- 三传 -->
  <view class="pan">
    <view class="row">
      <view class="item"><text>{{ shiPan.data.san_chuan.liu_qing[0] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.dun_gan[0] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.chu }}</text></view>
      <view class="item"><text>{{ toTianJiang(shiPan.data.san_chuan.chu) }}</text></view>
    </view>
    <view class="row">
      <view class="item"><text>{{ shiPan.data.san_chuan.liu_qing[1] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.dun_gan[1] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.zhong }}</text></view>
      <view class="item"><text>{{ toTianJiang(shiPan.data.san_chuan.zhong) }}</text></view>
    </view>
    <view class="row">
      <view class="item"><text>{{ shiPan.data.san_chuan.liu_qing[2] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.dun_gan[2] }}</text></view>
      <view class="item"><text>{{ shiPan.data.san_chuan.mo }}</text></view>
      <view class="item"><text>{{ toTianJiang(shiPan.data.san_chuan.mo) }}</text></view>
    </view>
</view>

<!-- 四课 -->
<view class="pan">
<view class="row">
        <view class="item"><text>{{ toTianJiang(shiPan.data.sike.zhi_ying) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.sike.zhi_yang) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.sike.gan_ying) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.sike.gan_yang) }}</text></view>
</view>
    <view class="row">
        <view class="item"><text>{{ shiPan.data.sike.zhi_ying }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.zhi_yang }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.gan_ying }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.gan_yang }}</text></view>
    </view>

    <view class="row">
        <view class="item"><text>{{ shiPan.data.sike.zhi_yang }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.zhi }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.gan_yang }}</text></view>
        <view class="item"><text>{{ shiPan.data.sike.gan }}</text></view>
    </view>
</view>
   <!-- 天盘 -->
<view class="pan">
   <view class="row">
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[5]) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[6]) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[7]) }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[8]) }}</text></view>
   </view>
   <view class="row">
        <view class="item"><text>{{ shiPan.data.tian_pan[5] }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[6] }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[7] }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[8] }}</text></view>
   </view>
   <view class="row">
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[4]) }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[4] }}</text></view>
        <view class="item"></view>
        <view class="item"></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[9] }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[9]) }}</text></view>
   </view>
   <view class="row">
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[3]) }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[3] }}</text></view>
        <view class="item"><text></text></view>
        <view class="item"><text></text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[10] }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[10]) }}</text></view>
   </view>
   <view class="row">
        <view class="item"><text>{{ shiPan.data.tian_pan[2]  }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[1]  }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[0]  }}</text></view>
        <view class="item"><text>{{ shiPan.data.tian_pan[11] }}</text></view>
   </view>
   <view class="row">
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[2])  }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[1])  }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[0])  }}</text></view>
        <view class="item"><text>{{ toTianJiang(shiPan.data.tian_pan[11]) }}</text></view>
   </view>
</view>
</template>
<script setup lang="ts">
import { diZhi } from "@/constdef";
import { daliurenAPI } from "@/http";
import type { DaliurenPan, RequestData } from "@/type";
import { onLoad } from "@dcloudio/uni-app";
import { reactive } from "vue";

let requestData: RequestData = reactive({
  year: 0,
  month: 1,
  day: 1,
  hour: 0,
  minute: 0,
  second: 0,
  yue_jiang: null,
  divination_time: "",
  diurnal: true,
  year_of_birth: 0,
  masculine: true,
});
const daliurenPan: DaliurenPan = {
  lunar_calendar: {
    year: "",
    month: "",
    day: "",
    time_gan_zhi: "",
  },
  solar_term_first: {
    name: "",
    year: 0,
    month: 0,
    day: 0,
    hour: 0,
    minute: 0,
    second: 0,
  },
  solar_term_second: {
    name: "",
    year: 0,
    month: 0,
    day: 0,
    hour: 0,
    minute: 0,
    second: 0,
  },
  si_zhu: {
    year: "",
    month: "",
    day: "",
    time: "",
  },
  kong: [],
  tian_pan: [],
  tian_jiang_pan: [],
  sike: {
    gan: "",
    gan_yang: "",
    gan_ying: "",
    zhi: "",
    zhi_yang: "",
    zhi_ying: "",
  },
  san_chuan: {
    chu: "",
    zhong: "",
    mo: "",
    dun_gan: [],
    liu_qing: [],
  },
  year_of_birth_gan_zhi: "",
  xing_nian: "",
  yue_jiang: "",
};

const shiPan = reactive({
  data: daliurenPan,
  completed: false,
  //   err: false,
  //   errMessage: "",
});

onLoad((option: any) => {
  const request = JSON.parse(decodeURIComponent(option.item));
  Object.assign(requestData, request);
  uni
    .request({
      url: daliurenAPI.url,
      method: daliurenAPI.method,
      header: daliurenAPI.header,
      data: request,
    })
    .then((res: any) => {
      shiPan.data = res.data;
      shiPan.completed = true;
    });
});

// 获取某地支的天将
function toTianJiang(d: string): string {
  let n = diZhi.indexOf(d);
  if (n == -1) return "";
  return shiPan.data.tian_jiang_pan[n];
}
</script>
<style lang="scss" scoped>
.info-content{
    display: flex;
    justify-content: center;
}
.info{
    display: flex;
    flex-direction: column;
    // justify-content: center;
    // align-content: center;
    .item{
        display: flex;
        margin: 5rpx;
    }
}

.pan{
    margin: 50rpx;
    .row{
        display: flex;
        justify-content: center;
        .item{
        width: 50rpx;
        }
    }
}
</style>