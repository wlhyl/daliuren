
export interface DaliurenPan {
    lunarCalendar: {
        year: string,
        month: string,
        day: string,
        timeGanZhi: string,
    },
    "solarTermFirst": {
        name: string,
        year: number,
        month: number,
        day: number,
        hour: number,
        minute: number,
        second: number,
    },
    "solarTermSecond": {
        name: string,
        year: number,
        month: number,
        day: number,
        hour: number,
        minute: number,
        second: number,
    },
    "ganZhi": {
        year: string,
        month: string,
        day: string,
        time: string,
    },
    "kongWang": string[],
    "tianPan": string[],
    "tianJiangPan": string[],
    "sike": {
        gan: string,
        ganYang: string,
        ganYing: string,
        zhi: string,
        zhiYang: string,
        zhiYing: string
    },
    "sanChuan": {
        chu: string,
        zhong: string,
        mo: string
        dunGan: string[],
        liuQing: string[],
    },
    yearOfBirthGanZhi: string,
    xingNian: string
}