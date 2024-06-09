using System;
namespace StudyBud.Data.Interfaces
{
	public interface IGeneralUserDAL
	{
		void NewUser(string userId, string firstName, string lastName, string email, string phone);
	}
}

