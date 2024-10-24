SERVER_CN=localhost

# Tạo CA private key với mã hóa DES3 và kích thước 4096-bit
openssl genrsa -passout pass:123456 -des3 -out ca.key 4096

# Tạo chứng chỉ CA tự ký (CA certificate) có thời hạn 10 năm
openssl req -passin pass:123456 -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/C=VN/ST=Hanoi/L=Hanoi/O=MyCompany/OU=IT Department/CN=$SERVER_CN"

# Tạo server private key với mã hóa DES3 và kích thước 4096-bit
openssl genrsa -passout pass:12345 -des3 -out server.key 4096

# Tạo yêu cầu chứng chỉ (CSR) cho server với SAN sử dụng tệp cấu hình
openssl req -passin pass:12345 -new -key server.key -out server.csr -config openssl-san.cnf

# Ký yêu cầu chứng chỉ của server (CSR) bằng CA certificate, tạo ra chứng chỉ server có thời hạn 10 năm
openssl x509 -req -passin pass:123456 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt -extensions req_ext -extfile openssl-san.cnf

# Chuyển server private key sang định dạng PKCS#8 không mã hóa
openssl pkcs8 -topk8 -nocrypt -passin pass:12345 -in server.key -out server.pem
