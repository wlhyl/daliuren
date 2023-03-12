export interface RequestData {
  year: number;
  month: number;
  day: number;
  hour: number;
  minute: number;
  second: number;
  yue_jiang: null | string;
  divination_time: string;
  diurnal: boolean;
  year_of_birth: number;
  masculine: boolean;
}

export interface DaliurenPan {
  lunar_calendar: {
    year: string;
    month: string;
    day: string;
    time_gan_zhi: string;
  };
  solar_term_first: SolarTerm;
  solar_term_second: SolarTerm;
  si_zhu: {
    year: string;
    month: string;
    day: string;
    time: string;
  };
  kong: string[];
  tian_pan: string[];
  tian_jiang_pan: string[];
  sike: {
    gan: string;
    gan_yang: string;
    gan_ying: string;
    zhi: string;
    zhi_yang: string;
    zhi_ying: string;
  };
  san_chuan: {
    chu: string;
    zhong: string;
    mo: string;
    dun_gan: string[];
    liu_qing: string[];
  };
  year_of_birth_gan_zhi: string;
  xing_nian: string;
  yue_jiang: string;
}

interface SolarTerm {
  name: string;
  year: number;
  month: number;
  day: number;
  hour: number;
  minute: number;
  second: number;
}
