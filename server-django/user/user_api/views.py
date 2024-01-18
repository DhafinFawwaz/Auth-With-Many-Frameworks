from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions
from .models import Mahasiswa, Todo
from .serializers import TodoSerializer, MahasiswaSerializer
from rest_framework_simplejwt.tokens import RefreshToken
from .custom_permission.is_not_authenticated import IsNotAuthenticated

class TodoListApiView(APIView):
    # add permission to check if user is authenticated
    permission_classes = [permissions.IsAuthenticated]

    # 1. List all
    def get(self, request, *args, **kwargs):
        
        todos = Todo.objects.filter(user = request.user.id)
        serializer = TodoSerializer(todos, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    # 2. Create
    def post(self, request, *args, **kwargs):

        print(request.user)

        data = {
            'task': request.data.get('task'), 
            'completed': request.data.get('completed'), 
            'mahasiswa_id': request.user.id
        }
        serializer = TodoSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
    

class RegisterApiView(APIView):
    permission_classes = [IsNotAuthenticated]
    
    # 2. Create
    def post(self, request, *args, **kwargs):
        try:
            new_user: Mahasiswa = Mahasiswa.objects.create_user(username = request.data.get('username'), email = request.data.get('email'), password = request.data.get('password'), nim = request.data.get('nim'))
            return Response(MahasiswaSerializer(new_user).data, status=status.HTTP_200_OK)
        except Exception as e:
            return Response({"message": str(e)}, status=status.HTTP_400_BAD_REQUEST)
    

class LoginApiView(APIView):
    permission_classes = [IsNotAuthenticated]
   
    # 1. List all for debugging, remove this later
    def get(self, request, *args, **kwargs):
        users = Mahasiswa.objects.all()
        return Response(users.values(), status=status.HTTP_200_OK)

    # 2. Login
    def post(self, request, *args, **kwargs):
        try:
            mahasiswa = Mahasiswa.objects.get(email = request.data.get('email'))
            if mahasiswa.check_password(request.data.get('password')):
                refresh = RefreshToken.for_user(mahasiswa)
                user_data = {
                    "username": mahasiswa.username,
                    "email": mahasiswa.email,
                    "password": mahasiswa.password,
                    "nim": mahasiswa.nim,
                    "refreshToken": str(refresh),
                    "accessToken": str(refresh.access_token),
                }
                return Response(user_data, status=status.HTTP_200_OK)
            else:
                return Response({"message": "Wrong password"}, status=status.HTTP_401_UNAUTHORIZED)
        except Mahasiswa.DoesNotExist:
            return Response({"message": "User not found"}, status=status.HTTP_404_NOT_FOUND)
        
class AuthApiView(APIView):
    permission_classes = [permissions.IsAuthenticated]

    def post(self, request, *args, **kwargs):
        try:
            mahasiswa = Mahasiswa.objects.get(pk = request.user.id) # request.user is decoded by auth
            return Response(MahasiswaSerializer(mahasiswa).data, status=status.HTTP_200_OK)
        except Mahasiswa.DoesNotExist:
            return Response({"message": "User not found"}, status=status.HTTP_404_NOT_FOUND)