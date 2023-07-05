# make sure you have ../sample and go mod init app there
go build -o ../sample/codegen
cd ../sample
./codegen init something --zap --fiber --wmsub
./codegen init something1 --zap --viper --wmsub --workjob
./codegen init something2 --zap --viper --gorm --fiber --workjob
./codegen init something3 --zap --viper --gorm --mongo
./codegen init something4 --zap --viper --gorm --mongo --wmpub
./codegen init something5 --zap --viper --gorm --mongo --wmpub --worken
cd ../code-generator
