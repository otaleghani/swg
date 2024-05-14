/******************************************************************

Author      Oliviero Taleghani
Date        2024-05-13

Description
Entry point for the package. Here you will have all of the functions
needed to generate random names. 

Usage
  func FirstName(lang string, format string) (string, error)
  | Gives a first name in the given language and format

  func LastName(lang string, format string) (string, error)
  | Gives a last name in the given language and format

  func FullName(lang string, format string, div string) (string, error)
  | Gives a full name in the given language and format

Dependency
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"

Todo

Changelog
  [0.0.1]   2024-05-13
  Added     Initial release

********************************************************************/

package spg

import (
  "github.com/otaleghani/spg/internal/parser"
  "github.com/otaleghani/spg/internal/formatter"
)

func SingleFirstName(lang string, format string) (string, error) {
  word, err := parser.ParseCsvSingleWord("names/" + lang + ".csv", 0)
  if err != nil {
    return "", err
  }
  err = formatter.Switch(&word, format)
  if err != nil {
    return "", err
  }
  return word, nil
}

func SingleLastName(lang string, format string) (string, error) {
  word, err := parser.ParseCsvSingleWord("surnames/" + lang + ".csv", 0)
  if err != nil {
    return "", err
  }
  err = formatter.Switch(&word, format)
  if err != nil {
    return "", err
  }
  return word, nil
}

func SingleFullName(lang string, format string, div string) (string, error) {
  var array []string
  firstName, err := SingleFirstName(lang, format)
  if err != nil {
    return "", err
  }
  lastName, err := SingleLastName(lang, format)
  if err != nil {
    return "", err
  }
  array = append(array, firstName, lastName)
  return formatter.Divider(array, div), nil
}

func SingleDomainName(format string) (string, error) {
  word, err := parser.ParseCsvSingleWord("domains/country.csv", 1)
  if err != nil {
    return "", err
  }
  err = formatter.Switch(&word, format)
  if err != nil {
    return "", err
  }
  return word, nil
}