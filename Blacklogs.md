# Make by KKGO
วิธีจัดการเงินด้วยแอปบริหารเงิน MAKE by KKGO ที่อยากจะแนะนำเชื่อว่าจะช่วยให้ทุกคนจบปัญหาใช้เงินเกินตัวหรือมีค่าใช้ 
จัดการเงินด้วยเงินจริงหนึ่งบัญชีแบ่งเป็นกระเป๋าย่อยได้ ตอบโจทย์ทุกไลฟ์สไตล์ตามใจคุณ
ใช้จ่ายโดยตรงจาก mobile banking บน MAKE ได้เลยทั้ง จัด จ่าย จด ทำทั้งหมดจบในที่เดียว
เปลี่ยนได้ตามใจ:เปลี่ยนชื่อและรูปด้วย Cloud Pocket ให้เข้ากับทุกเป้าหมายทั้งเล็กใหญ่
ใช้เสร็จก็ลบทิ้ง: Cloud Pocket ไหนไม่ใช้แล้ว สามารถลบทิ้งได้ ไม่มีการผูกมัดใดๆ
แชร์เก็บกับเพื่อน: ชวนเพื่อนมาจอย Cloud Pocket ร่วมกันเพื่อบริหารเงินหรือเก็บเงินกองกลางและบรรลุเป้าหมายไปด้วยกัน
จัด จ่าย จด: ทำทั้งหมดจบในที่เดียว MAKE ทำให้การใช้เงินในชีวิตประจำวันเป็นเรื่องง่าย ติดตามยอดใช้จ่ายสบาย ไม่ต้องคอยอัพเดตใน รายรับรายจ่ายแยกต่างหาก จึงมั่นใจได้ว่าเงินที่ จัดการอยู่เป็นจำนวนที่ตรงกับความเป็นจริง
Pop Pay: อยู่ด้วยกัน โอนให้กันได้ โอนเงินเเบบไร้สัมผัส (Contactless) ด้วย Pop Pay เมื่ออยู่ใกล้กันในรัศมี 10 เมตร ไอคอนของเพื่อนคุณจะปรากฏขึ้น สามารถกดเปย์ได้ทันที ไม่ง้อเลขบัญชีหรือเบอร์โทร!
Chat Banking: รายการที่ไม่ต้องตีความ หมดยุคของเลขบัญชี xxx-x-xxxxx-x ที่ไม่รู้ว่าของใคร หรือภาษาเเบ้งกิ้งชวนเวียนหัว เพราะ MAKE โชว์รายการในรูปแบบแชทเข้าใจง่าย แยกเป็นรายคนแถมเเนบโน้ตกับรูปได้ด้วย
Money Request: เรียกเก็บเงินคนที่ยืมไปได้ง่ายๆให้เพื่อนยืมเงินเมื่อไหร่ ก็ส่งรายการเรียกเก็บเงินคืนผ่าน MAKE ได้ คืนเงินสะดวก ไม่มีลืม
Expense Summary: เข้าใจการใช้จ่ายยิ่งกว่าที่เคย สรุปและติดตามการใช้จ่ายของคุณ ทำให้ทุกเดือนของคุณวางแผนการใช้จ่ายได้ดียิ่งขึ้น

## 1. Blacklogs



## "default cloud pocket" feature
### Story:
* As a user,
* I want my main account balance to be automatically added as a default "cloud pocket" called "Cashbox",
* so that I can easily access and manage my funds without having to manually create a new "cloud pocket" every time.

Acceptance Criteria:

- My main account balance is automatically added as a "cloud pocket" called "Cashbox" when I create an account or login.
- The balance of the "Cashbox" is always in sync with my main account balance
- I can access the "Cashbox" from the app's home screen
- I can see the balance and transaction history of "Cashbox" at all times
- I can still create additional "cloud pockets" if I want to budget or manage specific funds separately

* ในฐานะผู้ใช้,
* ฉันต้องการให้ยอดเงินหลักของฉันเป็น "cloud pocket" เริ่มต้นที่เรียกว่า "Cashbox",
* เพื่อให้ฉันสามารถเข้าถึงและจัดการเงินของฉันได้อย่างง่ายดายโดยไม่ต้องสร้าง "cloud pocket" ใหม่ในแต่ละครั้ง

เกณฑ์การยอมรับ:

- เมื่อฉันสร้างบัญชีหรือเข้าสู่ระบบยอดเงินหลักของฉันจะถูกเพิ่มเป็น "cloud pocket" เรียกว่า "Cashbox"
- ยอดเงิน "Cashbox" จะตรงกับยอดเงินหลักเสมอ
- ฉันสามารถเข้าถึง "Cashbox" จากหน้าแรกของแอป
- ฉันสามารถดูยอดเงินและประวัติการทำธุรกรรมของ "Cashbox" ได้ตลอด

## "cloud pocket" feature 
### Story:
* As a user,
* I want to be able to create multiple "cloud pockets" within my mobile banking app,
* so that I can budget and manage my money more effectively.

Acceptance Criteria:

- I can create a new "cloud pocket" from the app's home screen
- I can assign a name and budget to each "cloud pocket"
- I can add money to a "cloud pocket" from my main account balance
- I can view the current balance and transaction history for each "cloud pocket"
- I can transfer money between my "cloud pockets" and main account balance
				
* ในฐานะผู้ใช้,
* ฉันต้องการที่จะสามารถสร้าง "cloud pockets" หลายๆ อันในแอปธนาคารของฉันได้,
* เพื่อให้ฉันสามารถจัดการและจัดสรรเงินได้อย่างเหมาะสมยิ่งขึ้น

เกณฑ์การยอมรับ:

- ฉันสามารถสร้าง "cloud pockets" ใหม่จากหน้าแรกของแอป
- ฉันสามารถกำหนดชื่อและงบประมาณให้กับแต่ละ "cloud pockets"
- ฉันสามารถเพิ่มเงินไปยัง "cloud pockets" จากยอดเงินหลักของฉัน
- ฉันสามารถดูยอดเงินปัจจุบันและประวัติการทำธุรกรรมของแต่ละ "cloud pockets"
- ฉันสามารถโอนเงินระหว่าง "cloud pockets" และยอดเงินหลัก


## "syncing cloud pockets" feature
### Story:
* As a user,
* I want my "cloud pockets" to always be in sync with my main account balance,
* so that I always know how much money I have available no matter which "cloud pocket" I am viewing or moving money from.

Acceptance Criteria:

- The total balance of all "cloud pockets" is always equal to the balance of my main account
- Whenever I move money between "cloud pockets", the main account balance is updated accordingly
- I can always see the main account balance next to the balance of each "cloud pocket" for easy comparison
- If the main account balance changes( ex: deposit, withdrawal from other channel) the "cloud pockets" will be updated to reflect the change as well.

* ในฐานะผู้ใช้,
* ฉันต้องการให้ "cloud pockets" ของฉันอยู่ในสถานะตรงกับยอดเงินหลักของฉันเสมอ,
* เพื่อให้ฉันสามารถทราบว่าฉันมีเงินเหลือทั้งหมดกี่บาทได้ตลอดเวลาแม้ว่าฉันจะดูหรือโอนเงินจาก "cloud pocket" ไหน

เกณฑ์การยอมรับ:

- ยอดเงินของแต่ละ "cloud pockets" จะตรงกับยอดเงินหลักเสมอ
- เมื่อฉันโอนเงินระหว่าง "cloud pockets" ยอดเงินหลักจะถูกอัพเดทตามไปด้วย
- ฉันสามารถเห็นยอดเงินหลักได้ตลอดเวลาข้างหน้ายอดเงินของแต่ละ "cloud pocket" เพื่อเปรียบเทียบได้ง่าย
- เมื่อยอดเงินหลักมีการเปลี่ยนแปลง(เช่น เงินเข้าบัญชี, เงินถอนจากช่องทางอื่น) "cloud pockets" ก็จะถูกอัพเดทเป็นไปตาม


## "viewing cloud pockets" feature
### story:
* As a user,
* I want to be able to view a list of all my "cloud pockets" in one place,
* so that I can easily manage and track my funds.

Acceptance Criteria:

- I can access a list of all my "cloud pockets" from the app's home screen or main menu
- The list includes the name, current balance, and transaction history of each "cloud pocket"
- I can tap on a "cloud pocket" to view more details, such as recent transactions
- I can easily create new "cloud pockets" from this screen as well
- I can sort, filter or search my cloud pockets by name and balance

* ในฐานะผู้ใช้,
* ฉันต้องการสามารถดูรายการของ "cloud pockets" ทั้งหมดในที่เดียว,
* เพื่อให้ฉันสามารถจัดการและติดตามเงินของฉันได้อย่างง่ายดาย

เกณฑ์การยอมรับ:

- ฉันสามารถเข้าถึงรายการของ "cloud pockets" ทั้งหมดจากหน้าแรกหรือเมนูหลักของแอป
- รายการประกอบด้วยชื่อ, ยอดเงินปัจจุบัน และประวัติการทำธุรกรรมของแต่ละ "cloud pocket"
- ฉันสามารถคลิกที่ "cloud pocket" เพื่อดูรายละเอียดเพิ่มเติมเช่นประวัติการทำรายการล่าสุด
- ฉันสามารถสร้าง "cloud pockets" ใหม่ได้จากหน้านี้เช่นกัน
- ฉันสามารถเรียงลำดับ, กรอง, หรือค้นหา "cloud pockets" ของฉันโดยใช้ชื่อ และยอดเงินได้

## "cloud pockets transaction history" feature 
### story:
* As a user,
* I want to be able to view the transaction history of each of my "cloud pockets",
* so that I can keep track of my spending and budgeting.

Acceptance Criteria:

- I can view the transaction history of each "cloud pocket" by selecting it from the list of "cloud pockets"
- The transaction history includes the date, amount, and description of each transaction
- I can filter the transaction history by date range, transaction type or amount
- I can see the balance of each transaction and the remaining balance after it
- I can see the aggregate monthly balance and transaction history of the cloud pocket
- I can export the transaction history of each cloud pocket as .csv file

* ในฐานะผู้ใช้,
* ฉันต้องการสามารถดูประวัติการทำธุรกรรมของแต่ละ "cloud pockets" ของฉันได้,
* เพื่อให้ฉันสามารถติดตามการใช้จ่ายและการจัดงบประมาณได้

เกณฑ์การยอมรับ:

- ฉันสามารถดูประวัติการทำธุรกรรมของแต่ละ "cloud pocket" โดยเลือกจากรายการ "cloud pockets"
- ประวัติการทำธุรกรรมประกอบด้วยวันที่, จำนวนเงินและคำอธิบายของแต่ละรายการ
- ฉันสามารถกรองประวัติการทำธุรกรรมโดยใช้ช่วงเวลา, ประเภทธุรกรรม หรือ จำนวนเงิน
- ฉันสามารถเห็นยอดเงินคงเหลือของแต่ละรายการและยอดเงินคงเหลือหลังจากรายการ
- ฉันสามารถเห็นยอดเงินเดือนละและประวัติธุรกรรมของ cloud pocket
- ฉันสามารถ export ประวัติธุรกรรมของแต่ละ cloud pocket เป็นไฟล์ .csv

## Story:
* As a user, 
* I want to be able to transfer money between my cloud pockets within the app 
* so that I can easily manage my finances and budget my money in different categories.

Acceptance Criteria:

- I can select the source cloud pocket and the destination cloud pocket from a list of my existing cloud pockets.
- I can enter the amount of money I want to transfer.
- The app will show the remaining balance of each cloud pocket after the transfer.
- I can see a confirmation screen with the details of the transfer before it is completed.
- I can view the transaction history of each cloud pocket after the transfer to confirm the transaction has been made
- I can filter the transaction history by date, category, and amount

* ในฐานะผู้ใช้,
* ฉันต้องการสามารถโอนเงินระหว่าง cloud pockets ในแอปของฉันได้
* เพื่อให้ฉันสามารถจัดการเงินและงบประมาณเงินของฉันในหมวดหมู่ต่างๆได้ง่ายขึ้น

เกณฑ์การยอมรับ:

- ฉันสามารถเลือก cloud pocket ต้นทางและปลายทางจากรายการ cloud pockets ที่มีอยู่ในขณะนั้น
- ฉันสามารถป้อนจำนวนเงินที่ฉันต้องการโอน
- แอปจะแสดงยอดเงินคงเหลือของแต่ละ cloud pocket หลังจากโอน
- ฉันสามารถเห็นหน้าจอยืนยันพร้อมรายละเอียดการโอนก่อนที่จะทำการโอน
- ฉันสามารถดูประวัติการโอนของแต่ละ cloud pocket หลังจากโอนเพื่อยืนยันว่าการโอนเสร็จสมบูรณ์
- ฉันสามารถกรองประวัติการโอนโดยใช้วันที่, หมวดหมู่ และจำนวนเงิน

## "transfer money between cloud pockets" feature - error handling
### Story:
* As a user, 
* I want to make sure that I can only transfer money between cloud pockets if I have enough money in the source cloud pocket,
* so that I don't accidentally overdraw my account.

Acceptance Criteria:

- The app will check the balance of the source cloud pocket before allowing a transfer to take place.
- If the source cloud pocket does not have enough money for the transfer, the transfer will not be allowed and an error message will be displayed.
- If the source cloud pocket does have enough money for the transfer, the transfer will be allowed and the balance of both the source and destination cloud pockets will be updated.
- I can see a confirmation screen with the details of the transfer before it is completed.
- I can view the transaction history of each cloud pocket after the transfer to confirm the transaction has been made
- I can filter the transaction history by date, category, and amount

This story make sure that the user can't accidentally transfer money more than what they have in their account,and the error message will give clear indication to user what went wrong and how they can fix it.

* ในฐานะผู้ใช้,
* ฉันต้องการให้แน่ใจว่าฉันสามารถโอนเงินระหว่าง cloud pockets ได้เมื่อมีเงินใน cloud pocket ต้นทางเพียงพอ 
* เพื่อไม่ให้เกิดการถอนเงินเกินจำนวน

เกณฑ์การยอมรับ

- แอปจะตรวจสอบยอดเงินของ cloud pocket ต้นทางก่อนที่จะทำการโอน
- ถ้า cloud pocket ต้นทางไม่มีเงินเพียงพอสำหรับการโอน การโอนจะไม่ได้รับอนุญาตและจะแสดงข้อความแจ้งเตือน
- ถ้า cloud pocket ต้นทางมีเงินเพียงพอสำหรับการโอน การโอนจะได้รับอนุญาตและยอดเงินของทั้ง cloud pocket ต้นทางและปลายทางจะถูกปรับปรุง
- ฉันสามารถเห็นหน้าจอยืนยันพร้อมรายละเอียดการโอนก่อนที่จะทำการโอน
- ฉันสามารถดูประวัติการโอนของแต่ละ cloud pocket หลังจากโอนเพื่อยืนยันว่าการโอนเสร็จสมบูรณ์
- ฉันสามารถกรองประวัติการโอนโดยใช้วันที่, หมวดหมู่ และจำนวนเงิน

ความต้องการนี้เพื่อให้แอปตรวจสามารถตรวจสอบว่าผู้ใช้มีเงินเพียงพอใน cloud pocket ต้นทางก่อนที่จะทำการโอน เพื่อป้องกันการถอนเงินเกินจำนวนโดยอัตโนมัติ และสามารถแจ้งเตือนผู้ใช้ถ้าเกิดข้อผิดพลาดในการโอนเงินได้

## Delete Cloud Pocket feature
### Story:
* As a user,
* I want to be able to delete a "cloud pocket" within my mobile banking app,
* so that I can manage my money more effectively and remove pockets that I no longer need.

Acceptance Criteria:

- I can delete a "cloud pocket" from the app's home screen
- I can only delete a "cloud pocket" if its balance is zero, if there's money in it, the system should prompt a message saying that it can't be deleted.
- After the deletion, I can view the transaction history of the deleted cloud pocket for a certain period of time (e.g. 30 days)
- I can still access the transaction history of the deleted cloud pocket by using a transaction history endpoint.
- After the certain period of time, the transaction history of the deleted cloud pocket will be permanently deleted from the system as well

* ในฐานะผู้ใช้,
* ฉันต้องการที่จะสามารถลบ "กระเป๋าเงินในคลาวด์" ในแอพมือถือธนาคารของฉันได้
* เพื่อให้ฉันสามารถจัดการเงินของฉันได้อย่างมีประสิทธิภาพและลบกระเป๋าที่ฉันไม่ได้ใช้อีกต่อไป

เงื่อนไขที่ต้องปฏิบัติ:

- ฉันสามารถลบ "กระเป๋าเงินในคลาวด์" จากหน้าแรกของแอพ
- ฉันสามารถลบ "กระเป๋าเงินในคลาวด์" ได้เมื่อเงินในกระเป๋านั้นเป็นศูนย์ หากมีเงินในนั้น ระบบจะแจ้งเตือนว่าไม่สามารถลบได้
- หลังจากการลบ ฉันยังสามารถดูประวัติการโอนเงินของกระเป๋าเงินที่ถูกลบ ภายในช่วงเวลาที่กำหนด (เช่น 30 วัน)
- ฉันสามารถเข้าถึงประวัติการโอนเงินของกระเป๋าเงินที่ถูกลบ โดยใช้ endpoint transaction
- หลังจากช่วงเวลาที่กำหนด ประวัติการโอนเงินของกระเป๋าเงินที่ถูกลบจะถูกลบออกจากระบบแบบถาวร

## Aunthentication and Authorization feature
### Story:
* As a user,
* I want to be able to securely authenticate and authorize my actions when using the Cloud Pocket system,
* so that I am able to access and make changes only to the resources that I am authorized to.

Acceptance Criteria:

- All API calls to the Cloud Pocket system must be authenticated and authorized before they are processed.
- There should be a registration/create new user API that allows users to create new accounts with the system.
- The registration API should require the user to provide a unique username and password, and the password should be securely hashed and stored.
- The authentication and authorization process should use a secure method such as OAuth 2.0 or JWT.
- The system should have a mechanism for revoking authentication and authorization for a user if necessary.
- The system should be able to check for duplicate/existing username on registration process
- The user should be able to change their password if they forget it by using the password recovery feature, it should also require a email verification.

This story ensure that, only authorized users are able to access and make changes to the resources within the Cloud Pocket system, and that the system has appropriate measures in place to ensure the security of user's credentials, account creation, and access.

* ในฐานะผู้ใช้,
* ฉันต้องการที่จะสามารถเข้ารหัสและอนุมัติการเปลี่ยนแปลงของฉันเมื่อใช้ระบบกระเป๋าเคลื่อนที่ได้อย่างปลอดภัย
* เพื่อให้ฉันสามารถเข้าถึงและเปลี่ยนแปลงได้เฉพาะที่ฉันมีสิทธิ์อนุญาต

เกณฑ์การยอมรับ:

- การเรียก API เข้าถึงระบบกระเป๋าเคลื่อนที่ต้องได้รับการตรวจสอบและอนุมัติก่อนที่จะประมวลผล
- มี API สำหรับลงทะเบียน/สร้างผู้ใช้ใหม่ โดยที่ผู้ใช้สามารถสร้างบัญชีใหม่กับระบบได้
- การลงทะเบียน API จะต้องการให้ผู้ใช้ให้ชื่อผู้ใช้และรหัสผ่านที่ไม่ซ้ำกัน และรหัสผ่านจะถูกเข้ารหัสและเก็บไว้อย่างปลอดภัย
- ระบบควรมีการตรวจสอบอีเมลล์ของผู้ใช้ในขั้นตอนการกู้คืนรหัสผ่าน

ความต้องการนี้ทำให้แน่ใจว่าเฉพาะผู้ใช้ที่ได้รับอนุญาตเข้าถึงและเปลี่ยนแปลงทรัพยากรในระบบกระเป๋าเคลื่อนที่ได้ และว่าระบบมีเกณฑ์ป้องกันความปลอดภัยของข้อมูลประจำตัวของผู้ใช้และการสร้างบัญชีและเข้าถึงได้


## Technical guide
- GET /cloud-pockets: Retrieve a list of all cloud pockets belonging to the authenticated user.
- POST /cloud-pockets: Create a new cloud pocket. The request body should contain the name, currency and initial balance of the cloud pocket.
- GET /cloud-pockets/:id: Retrieve the details of a specific cloud pocket, identified by its ID.
- PUT /cloud-pockets/:id: Update the details of a specific cloud pocket, identified by its ID.
- DELETE /cloud-pockets/:id: Delete a specific cloud pocket, identified by its ID.
- GET /cloud-pockets/:id/transactions: Retrieve the transaction history of a specific cloud pocket.
- POST /cloud-pockets/transfer: Transfer funds from one cloud pocket to another. The request body should contain the source cloud pocket ID, destination cloud pocket ID, amount and currency.
- GET /cloud-pockets/:id/balance: Retrieve the current balance of a specific cloud pocket,
- GET /cloud-pockets/:id/monthly-balance: Retrieve the aggregate monthly balance of a specific cloud pocket.
- GET /cloud-pockets/:id/csv: export transaction history of a specific cloud pocket as .csv file

- GET /cloud-pockets

Response Body

```json
[
		{
			"id": "12345",
			"name": "Travel Fund",
			"category": "Vacation",
			"currency": "USD",
			"balance": 100
		},
		{
			"id": "67890",
			"name": "Savings",
			"category": "Emergency Fund",
			"currency": "USD",
			"balance": 200
		}
	]
```

- POST /cloud-pockets

Request Body

```json
	{
			"name": "Travel Fund",
			"currency": "USD",
			"initial_balance": 100.00
	}
```

Reponse Body

```json
	{
			"id": "246810",
			"name": "Travel Fund",
			"category": "Vacation",
			"currency": "USD",
			"balance": 100.00
	}
```

- GET /cloud-pockets/:id 
Reponse Body

```json
	{
			"id": "12345",
			"name": "Travel Fund",
			"category": "Vacation",
			"currency": "USD",
			"balance": 100.00
	}
```

- PUT /cloud-pockets/:id

Request Body

```json
	{
			"name": "Holiday Fund",
			"currency": "USD",
			"initial_balance": 150.00
	}
```

Reponse Body

```json
	{
			"id": "12345",
			"name": "Holiday Fund",
			"category": "Vacation",
			"currency": "USD",
			"balance": 150.00
	}
```

- DELETE /cloud-pockets/:id:

```json
	{
			"message": "Cloud pocket deleted successfully"
	}
```

```json
{
    "message":"Unable to delete this Cloud Pocket\n there is amount left in this Cloud Pocket, please move money out and try again"
}
```

404 Not Found

```json
	{
			"message": "Cloud pocket not found"
	}
```


- GET /cloud-pockets/:id/transactions: This endpoint does not require a request body.

Reponse Body

```json
	{
			"id": "12345",
			"transactions": [
					{
							"date": "2022-01-01",
							"description": "Hotel booking",
							"amount": 50.00,
							"type": "debit"
					},
					{
							"date": "2022-01-02",
							"description": "Airfare",
							"amount": 200.00,
							"type": "debit"
					}
			]
	}
```

- POST /cloud-pockets/transfer:

Request body:

```json
	{
			"source_cloud_pocket_id": "12345",
			"destination_cloud_pocket_id": "67890",
			"amount": 50.00,
			"description":"Transfer from Travel fund to savings"
	}
```

Response body :

```json
	{
			"transaction_id": "123",
			"source_cloud_pocket": {
					"id": "12345",
					"name": "Travel Fund",
					"category": "Vacation",
					"balance": 50.00
			},
			"destination_cloud_pocket": {
					"id": "67890",
					"name": "Savings",
					"category": "Emergency Fund",
					"balance": 250.00
			},
			"status": "Success"
	}
```

or

```json
	{
			"error_message":"Not enough balance in the source cloud pocket",
			"status":"Failed"
	}
```

- GET /cloud-pockets/:id/transactions :

Request body: None

Response body:

```json
	{
			"transactions": [
					{
							"id": "123",
							"source_cloud_pocket_id":"12345",
							"destination_cloud_pocket_id":"67890",
							"amount": 50.00,
							"description":"Transfer from Travel fund to savings",
							"date":"2022-01-01",
							"status":"Success"
					},
					{
							"id": "124",
							"source_cloud_pocket_id":"67890",
							"destination_cloud_pocket_id":"12345",
							"amount": 25.00,
							"description":"Transfer from Savings to Travel fund",
							"date":"2022-01-02",
							"status":"Failed"
					}
			]
	}
```

- GET /cloud-pockets/:id/balance: 

Reponse Body

```json
	{
			"id": "12345",
			"balance": 100.00,
			"currency": "USD"
	}
```

-GET /cloud-pockets/:id/monthly-balance

Reponse Body

```json
	{
			"id": "12345",
			"monthly_balance": [
					{
							"month": "January",
							"year": "2022",
							"balance": 250.00
					},
					{
							"month": "February",
							"year": "2022",
							"balance": 200.00
					}
			]
	}
```

- GET /cloud-pockets/:id/csv: This endpoint does not require a request body.

- POST /cloud-pockets/transfer:
Request body:

```json
	{
			"source_cloud_pocket_id": "12345",
			"destination_cloud_pocket_id": "67890",
			"amount": 50.00,
			"description":"Transfer from Travel fund to savings"
	}
```

Response body:

```json
	{
			"transaction_id": "123",
			"source_cloud_pocket": {
					"id": "12345",
					"name": "Travel Fund",
					"category": "Vacation",
					"balance": 50.00
			},
			"destination_cloud_pocket": {
					"id": "67890",
					"name": "Savings",
					"category": "Emergency Fund",
					"balance": 250.00
			}
	}
```
- GET /cloud-pockets/:id/transactions:
Request body: None
Response body:

```json
	{
			"transactions": [
					{
							"id": "123",
							"source_cloud_pocket_id":"12345",
							"destination_cloud_pocket_id":"67890",
							"amount": 50.00,
							"description":"Transfer from Travel fund to savings",
							"date":"2022-01-01"
					},
					{
							"id": "124",
							"source_cloud_pocket_id":"67890",
							"destination_cloud_pocket_id":"12345",
							"amount": 25.00,
							"description":"Transfer from Savings to Travel fund",
							"date":"2022-01-02"
					}
			]
	}
```

### hotfix