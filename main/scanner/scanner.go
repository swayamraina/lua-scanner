package scanner


// @author 	swayam.raina
// @dated 	05.12.2019





// This structure defines the main scanner
// service which does the actual heavy
// lifting of file-system scanning
type Scanner struct {
	config Config
}

// This method does an internal and external scan
// to discover lua scripts in the project
func (scanner *Scanner) Scan () []string {
	var files []string
	config := scanner.config
	//if config.InternalScan {
	//	files = append(files, internal.Scan())
	//}
	//if config.ExternalScan {
	//	files = append(files, external.Scan())
	//}
	return files
}