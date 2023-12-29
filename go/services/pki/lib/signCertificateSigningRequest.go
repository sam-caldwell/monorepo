package pkicore

//
//var caPrivateKey *ecdsa.PrivateKey
//var caCertificate *x509.Certificate
//
//func init() {
//	const caCertFile = "/tmp/adrestia-pki.ca.cert"
//
//}
//
//func SignCertificateSigningRequest(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "Expects HTTP/POST", http.StatusMethodNotAllowed)
//		return
//	}
//	// Read the uploaded CSR file
//	csrFile, _, err := r.FormFile("csr")
//	if err != nil {
//		http.Error(w, "Failed to read CSR file", http.StatusBadRequest)
//		return
//	}
//	defer func() { _ = csrFile.Close() }()
//
//	csrBytes, err := io.ReadAll(csrFile)
//	if err != nil {
//		http.Error(w, "Failed to read CSR file content", http.StatusInternalServerError)
//		return
//	}
//
//	// Parse CSR
//	csr, _ := x509.ParseCertificateRequest(csrBytes)
//
//	// Sign the CSR with the CA private key
//	certDER, _ := x509.CreateCertificate(rand.Reader, csr, caCertificate, &caPrivateKey.PublicKey, caPrivateKey)
//
//	// Write the signed certificate to the response
//	w.Header().Set("Content-Type", "application/x-x509-ca-cert")
//	w.Header().Set("Content-Disposition", "attachment; filename=signed_cert.crt")
//	if _, err := w.Write(certDER); err != nil {
//		log.Fatalf("Error writing response: %v", err)
//	}
//}
