targetScope = 'subscription'

param location string = deployment().location

resource studyGroup 'Microsoft.Resources/resourceGroups@2021-04-01' = {
  name: 'rg-study'
  location: location
}

module storage './storage.bicep' = {
  name: 'storage'
  scope: studyGroup
  params: {
    storageAccountName: 'ststudy'
    location: location
  }
}
