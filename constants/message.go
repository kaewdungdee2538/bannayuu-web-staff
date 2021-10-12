package constants

const MessageSuccess = "เรียบร้อย"
const MessageFailed = "ทำรายการล้มเหลว"

const MessageCombineFailed = "Combine Error"
const MessageDataNotCompletely = "กรอกข้อมูลไม่ครบถ้วน"

const MessageNotAuthorization = "Session หมดอายุ"
const MessageAuthorizationBearerNotSet = "Bearer Token is null!"
const MessageUsernameOrPasswordNotValid = "Username หรือ Password ไม่ถูกต้อง"
const MessageUsernameNotFount = "กรุณากรอก Username"
const MessagePasswordNotFount = "กรุณากรอก Password"
const MessageUsernameIsSpecialProhibit = "Username ต้องเป็นอักษรภาษอังกฤษ หรือตัวเลขเท่านั้น"
const MessagePasswordIsSpecialProhibit = "Password ต้องเป็นอักษรภาษอังกฤษ หรือตัวเลขเท่านั้น"
const MessageUserIsDuplicate = "Username ซ้ำในระบบ"

const MessageCompanyIdNotFound = "ไม่พบรหัสโครงการ"
const MessageCompanyIdNotNumber = "รหัสโครงการต้องเป็นตัวเลขเท่านั้น"
const MessageCompanyCodeNotFount = "กรุณากรอกรหัสโครงการ"
const MessageCompanyCodeIsSpecialProhibit = "รหัสโครงการ ต้องเป็นอักษรภาษอังกฤษ หรือตัวเลขเท่านั้น"
const MessageCompanyNameNotFount = "กรุณากรอกชื่อโครงการ"
const MessageCompanyNameIsSpecialProhibit = "ชื่อโครงการ ห้ามมีอักขระพิเศษ"
const MessageCompanyProNotFound = "กรุณาเลือก Pro"
const MessageCompanyProIsSpecialProhibit = "Pro ต้องเป็นอักษรภาษอังกฤษ หรือตัวเลขเท่านั้น"
const MessageCompanyNotInBase = "ไม่พบโครงการในระบบ"
const MessageCompanyIsDuplicateInBase = "รหัสโครงการ หรือชื่อโครงการซ้ำในระบบ"
const MessageOldCompanyIdNotFound = "ไม่พบรหัสโครงการปัจจุบัน"
const MessageOldCompanyIdNotNumber = "รหัสโครงการปัจจุบัน ต้องเป็นตัวเลขเท่านั้น"
const MessageNewCompanyIdNotFound = "ไม่พบรหัสโครงการใหม่"
const MessageNewCompanyIdNotNumber = "รหัสโครงการใหม่ ต้องเป็นตัวเลขเท่านั้น"
const MessageOldCompanyNotInBase = "ไม่พบโครงการปัจจุบันในระบบ"
const MessageNewCompanyNotInBase = "ไม่พบโครงการใหม่ในระบบ"

const MessageDateStartNotFound = "กรุณากรอกเวลาเริ่มต้น"
const MessageDateStarFormatNotValid = "รูปแบบเวลาเริ่มต้นไม่ถูกต้อง"
const MessageDateStopNotFound = "กรุณากรอกเวลาสิ้นสุด"
const MessageDateStopFormatNotValid = "รูปแบบเวลาสิ้นสุดไม่ถูกต้อง"
const MessageCovertObjTOJSONFailed = "Convert To JSON Data Failed"

const MessageRemarkNotFount = "กรุณากรอกเหตุผล"
const MessageRemarkIsLower10Character = "กรุณากรอกเหตุผลมากกว่า 10 ตัวอักษร"
const MessageRemarkProhibitSpecial = "เหตุผล ห้ามมีอักขระพิเศษ"

const MessageImageNotFound = "ไม่พบรูปภาพ"

const MessageHomeAddressNotFound = "กรุณากรอกบ้านเลขที่"
const MessageHomeAddressProhibitSpecial = "บ้านเลขที่ ห้ามมีอักขระพิเศษ"

const MessageFullNameProhitbitSpecial = "ชื่อหรือนามสกุล ห้ามมีอักขระพิเศษ"
const MessageFirstNameNotFound = "กรุณากรอกชื่อ"
const MessageFirstNameProhitbitSpecial = "ชื่อ ห้ามมีอักขระพิเศษ"
const MessageLastNameNotFound = "กรุณากรอกนามสกุล"
const MessageLastNameProhitbitSpecial = "นามสกุล ห้ามมีมีอักขระพิเศษ"
const MessageTelNumberNotFound = "กรุณากรอกเบอร์โทรศัพท์"
const MessageTelNumberIsLess10Digit = "เบอร์โทรศัพท์ต้องไม่น้อยกว่า 10 ตัวอักษร"
const MessageTelNumberIsMoreThan10Digit = "เบอร์โทรศัพท์ต้องไม่เกิน 10 ตัวอักษร"
const MessageTelNumberNotNumber = "เบอรโทรศัพท์ต้องเป็นตัวเลขเท่านั้น"

const MessageAddressProhibitSpecial = "ที่อยู่ ห้ามมีอักขระพิเศษ"
const MessageLineProhibitSpecial = "ไลน์ ห้ามมีอักขระพิเศษ"
const MessageMobileNotNumber = "เบอร์โทรศัพท์ต้องเป็นตัวเลขเท่านั้น"
const MessageMobileNotEqual10Character = "เบอร์โทรศัพท์ต้องมี 10 ตัวอักษร"
const MessageEmailFormatInValid = "รูปแบบอีเมลไม่ถูกต้อง"
const MessageEmployeePrivilegeIdNotFound = "ไม่พบรหัสของสิทธิ์การเข้าใช้งานของ User"
const MessageEmployeePrivilegeIdNotNumber = "รหัสของสิทธิ์การเข้าใช้งานของ User ต้องเป็นตัวเลขเท่านั้น"
const MessageEmployeeStatusNotFound = "ไม่พบสถานะของ User"
const MessageEmployeeStatusProhibitSpecial = "สถานะของ User ห้ามมีอักขระพิเศษ"
const MessageEmployeeTypeNotFound = "ไม่พบประเภทของ User"
const MessageEmployeeTypeProhibitSpecial = "ประเภทของ User ห้ามมีอักขระพิเศษ"
const MessageEmployeeIdNotFound = "ไม่พบรหัสพนักงาน"
const MessageEmployeeIdNotNumber = "รหัสพนักงานต้องเป็นตัวเลขเท่านั้น"
const MessageUserNotInBase = "ไม่พบพนักงานในระบบ"
const MessageUserPrivilegeNotInBase = "ไม่พบสิทธิ์การเข้าใช้งานในระบบ หรือสิทธิ์ที่เลือกไม่ได้รับอนุญาติให้ใช้งาน"

const MessageHoldTimeNotFound = "ไม่พบเวลาที่กำหนด"
const MessageHoldTimeIsProhibitSpecial = "เวลาที่กำหนดห้ามมีอักษรพิเศษ ดังนี้ " + `/!@#$%^&*()_+\-=\[\]{};':"|,.<>\/?~`
