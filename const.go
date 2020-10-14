package sevenz

const sevenCmd = "7z"
const extractCmd = "x"
const archiveCmd = "a"
const destFolderTemplate = "-o{folder}"
const templatePattern = "{folder}"
const insaneCompressionParams = "-mx=9 -mfb=273 -ms -md=31 -myx=9 -mtm=- -mmt -mmtf -md=1536m -mmf=bt3 -mmc=10000 -mpb=0 -mlc=0"
const archiveType = "-t7z"

/*
7z a -t7z -mx=9 -mfb=273 -ms -md=31 -myx=9 -mtm=- -mmt -mmtf -md=1536m -mmf=bt3 -mmc=10000 -mpb=0 -mlc=0 archive.7z
insane compression
7z a -t7z -mx=9 -mfb=273 -ms -md=31 -myx=9 -mtm=- -mmt -mmtf -md=1536m -mmf=bt3 -mmc=10000 -mpb=0 -mlc=0
  archive.7z inputfileordir
*/
