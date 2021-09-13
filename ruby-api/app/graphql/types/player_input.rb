module Types
  class PlayerInput < BaseInputObject
    argument :age,  Integer, required: true
    argument :name, String,  required: true
  end
end
