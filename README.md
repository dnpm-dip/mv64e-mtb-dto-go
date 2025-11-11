# DNPM:DIP MTB Data-Transfer-Objects for Golang

Serialization and deserialization of DNPM:DIP MTB DTOs for the Go programming language.

This crate provides DNPM:DIP data model for use with "Modellvorhaben gem. §64e SGB V"

**Warning**: As of now, this package includes no date format conversion in JSON serialization
for patients birthdate and date of death related to 
https://github.com/dnpm-dip/backend-core/commit/97c44aa8bbd6ba4ac81824c5178db23fd08f9068
as has been included in https://github.com/dnpm-dip/mv64e-mtb-dto-java and
https://github.com/dnpm-dip/mv64e-mtb-dto-dotnet

