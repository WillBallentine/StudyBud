using System;
namespace StudyBud.Data.Interfaces
{
	public interface IGeneralUserDAL
	{
		void NewUserAsync(string Id, string phone, string fName, string lName);
	}
}

