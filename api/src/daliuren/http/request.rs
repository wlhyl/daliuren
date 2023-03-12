use serde::Deserialize;
use validator::Validate;

#[derive(Deserialize, Validate)]
pub struct DaLiuRenReust {
    #[validate(range(min = 1900))]
    pub year: i32,
    #[validate(range(min = 1, max = 12))]
    pub month: u8,
    #[validate(range(min = 1, max = 31))]
    pub day: u8,
    #[validate(range(min = 0, max = 23))]
    pub hour: u8,
    #[validate(range(min = 0, max = 59))]
    pub minute: u8,
    #[validate(range(min = 0, max = 59))]
    pub second: u8,
    #[validate(length(min = 1), non_control_character)]
    pub yue_jiang: Option<String>, // 月将，如：寅，空，表示自动计算
    #[validate(length(min = 1), non_control_character)]
    pub divination_time: String, // 占时，如：寅
    pub diurnal: bool, // 昼占:true，夜占:false
    #[validate(range(min = 1900))]
    pub year_of_birth: i32, // 出生年
    pub masculine: bool, // 性别，男：true，女：false
}
