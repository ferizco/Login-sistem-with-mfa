# Apa itu MFA? 
### MFA (Multi-Factor Authentication) adalah metode keamanan yang memerlukan lebih dari satu bentuk verifikasi untuk mengonfirmasi identitas pengguna ketika mencoba mengakses sistem atau aplikasi.

# Apa saja faktornya?
- Faktor pengetahuan: Sesuatu yang hanya diketahui pengguna, seperti kata sandi atau PIN.
- Faktor kepemilikan: Sesuatu yang dimiliki pengguna, seperti perangkat keras (token fisik, ponsel, atau aplikasi autentikasi).
- Faktor inheren: Sesuatu yang dimiliki pengguna secara biologis, seperti sidik jari, pengenalan wajah, atau pemindaian iris mata.

### Membuat MFA Menggunakan OTP Based dengan Google Authenticator dan Golang. 

#  Alur Aplikasi 
1. Pengguna melakukan signup 
2. qr code akan muncul, lalu pengguna melakukan scan qr code tersebut di google authenticator 
3. pengguna melakukan login menggunakan username dan password 
4. jika valid, pengguna harus memasukkan juga code otp yang ada di google authenticator
5. jika valid baru pengguna akan bisa login. 